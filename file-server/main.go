package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const (
	fileStorageEnv = "FILE_STORAGE"
	httpsPortEnv   = "HTTPS_PORT"
	httpPortEnv    = "HTTP_PORT"
	certPathEnv    = "CERT_PATH"
	keyPathEnv     = "KEY_PATH"
)

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
	log.Printf("https server start on port %s with config: %+v", conf.HttpsPort(), conf)
	log.Fatal(httpsServer.ListenAndServeTLS(conf.CertPath, conf.KeyPath))
}

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
