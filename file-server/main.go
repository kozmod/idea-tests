package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

const (
	fileStorageEnv = "FILE_STORAGE"
	httpsPortEnv   = "HTTPS_PORT"
	httpPortEnv    = "HTTP_PORT"
	certPathEnv    = "CERT_PATH"
	keyPathEnv     = "KEY_PATH"
)

//cp sec.syn.crt  /usr/local/share/ca-certificates/sec.syn.crt && update-ca-certificates
//openssl x509 -in ott-ingress.crt -text

func handleTunneling(w http.ResponseWriter, r *http.Request) {
	dest_conn, err := net.DialTimeout("tcp", "localhost:9003", 10*time.Second) //todo: fix
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}
	client_conn, _, err := hijacker.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
	}
	go transfer(dest_conn, client_conn)
	go transfer(client_conn, dest_conn)
}
func transfer(destination io.WriteCloser, source io.ReadCloser) {
	defer destination.Close()
	defer source.Close()
	io.Copy(destination, source)
}
func handleHTTP(w http.ResponseWriter, req *http.Request) {
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()
	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func main() {
	conf, err := configFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	httpMux := http.NewServeMux()
	httpMux.HandleFunc("/status", status(conf))
	httpMux.HandleFunc("/", files(conf.FilesDir))

	httpServer := http.Server{
		Addr:    conf.HttpPort(),
		Handler: httpMux,
	}
	go func() {
		log.Printf("http server start on port %s with config: %+v", conf.HttpPort(), conf)
		log.Fatal(httpServer.ListenAndServe())
	}()

	httpsMux := http.NewServeMux()
	httpsMux.HandleFunc("/status", status(conf))
	httpsMux.HandleFunc("/", files(conf.FilesDir))
	httpsServer := http.Server{
		Addr:    conf.HttpsPort(),
		Handler: httpMux,
	}

	go func() {
		log.Printf("https server start on port %s with config: %+v", conf.HttpsPort(), conf)
		log.Fatal(httpsServer.ListenAndServeTLS(conf.CertPath, conf.KeyPath))
	}()

	server := &http.Server{
		Addr: ":8888",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodConnect {
				handleTunneling(w, r)
			} else {
				handleHTTP(w, r)
			}
		}),
		// Disable HTTP/2.
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}
	log.Fatal(server.ListenAndServe())
}

//func main() {
//	conf, err := configFromEnv()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	httpMux := http.NewServeMux()
//	httpMux.HandleFunc("/status", status(conf))
//	httpMux.HandleFunc("/", files(conf.FilesDir))
//
//	httpServer := http.Server{
//		Addr:    conf.HttpPort(),
//		Handler: httpMux,
//	}
//	go func() {
//		log.Printf("http server start on port %s with config: %+v", conf.HttpPort(), conf)
//		log.Fatal(httpServer.ListenAndServe())
//	}()
//
//	httpsMux := http.NewServeMux()
//	httpsMux.HandleFunc("/status", status(conf))
//	httpsMux.HandleFunc("/", files(conf.FilesDir))
//	httpsServer := http.Server{
//		Addr:    conf.HttpsPort(),
//		Handler: httpMux,
//	}
//	log.Printf("https server start on port %s with config: %+v", conf.HttpsPort(), conf)
//	log.Fatal(httpsServer.ListenAndServeTLS(conf.CertPath, conf.KeyPath))
//}
//
func status(cong Conf) func(writer http.ResponseWriter, request *http.Request) {
	confByte, err := json.Marshal(cong)
	if err != nil {
		log.Fatal(err)
	}
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "text/plain")

		if _, err := writer.Write(confByte); err != nil {
			log.Print(err)
		}
	}
}

func files(dir string) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		log.Println(request.URL)
		spath := strings.Split(request.URL.Path, "/")
		file := spath[len(spath)-1]
		filePath := fmt.Sprintf("%s/%s", dir, file)
		log.Printf("url path: %s, file path: %s", request.URL.Path, filePath)
		data, err := ioutil.ReadFile(filePath)
		if err != nil {
			_, _ = writer.Write([]byte(fmt.Sprintf("file not found - url path: %s, file path: %s", request.URL.Path, filePath)))
		} else {
			if _, err := writer.Write(data); err != nil {
				_, _ = writer.Write([]byte(fmt.Sprintf("problem reading file - url path: %s, file path: %s", request.URL.Path, filePath)))
			}
		}
	}
}

type Conf struct {
	FilesDir  string
	CertPath  string
	KeyPath   string
	httpPort  uint64
	httpsPort uint64
}

func (c Conf) HttpsPort() string {
	return ":" + strconv.FormatUint(c.httpsPort, 10)
}

func (c Conf) HttpPort() string {
	return ":" + strconv.FormatUint(c.httpPort, 10)
}

func configFromEnv() (Conf, error) {
	conf := Conf{}
	conf.FilesDir = os.Getenv(fileStorageEnv)
	httpsPort, err := strconv.ParseUint(os.Getenv(httpsPortEnv), 10, 32)
	if err != nil {
		return conf, errors.WithMessage(err, fmt.Sprintf("parce env %s", httpsPortEnv))
	}
	conf.httpsPort = httpsPort
	httpPort, err := strconv.ParseUint(os.Getenv(httpPortEnv), 10, 32)
	if err != nil {
		return conf, errors.WithMessage(err, fmt.Sprintf("parce env %s", httpPortEnv))
	}
	conf.httpPort = httpPort
	conf.CertPath = os.Getenv(certPathEnv)
	conf.KeyPath = os.Getenv(keyPathEnv)
	return conf, nil
}
