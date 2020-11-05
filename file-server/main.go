package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const (
	fileStorageEnv = "FILE_STORAGE"
	portEnv        = "PORT"
	certPathEnv    = "CERT_PATH"
	keyPathEnv     = "KEY_PATH"
)

func main() {
	conf, err := configFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	//pkg.GenCerts(
	//	pkg.WithHost("127.0.0.1"+conf.Port()),
	//	pkg.WithCertAndKeyPaths(conf.CertPath, conf.KeyPath),
	//)

	http.HandleFunc("/status", func(writer http.ResponseWriter, request *http.Request) {
		confByte, err := json.Marshal(conf)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			_, _ = writer.Write([]byte(err.Error()))

			return
		}
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "text/plain")

		if _, err := writer.Write(confByte); err != nil {
			log.Print(err)
		}
	})

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.URL)
		spath := strings.Split(request.URL.Path, "/")
		file := spath[len(spath)-1]
		filePath := fmt.Sprintf("%s/%s", conf.FilesDir, file)
		log.Printf("url path: %s, file path: %s", request.URL.Path, filePath)
		http.ServeFile(writer, request, filePath)
	})

	log.Printf("server start on port %s  with config: %+v", conf.Port(), conf)
	//log.Fatal(http.ListenAndServe(conf.Port(), nil))
	log.Fatal(http.ListenAndServeTLS(conf.Port(), conf.CertPath, conf.KeyPath, nil))
}

type Conf struct {
	FilesDir string
	CertPath string
	KeyPath  string
	port     uint64
}

func (c Conf) Port() string {
	return ":" + strconv.FormatUint(c.port, 10)
}

func configFromEnv() (Conf, error) {
	conf := Conf{}
	conf.FilesDir = os.Getenv(fileStorageEnv)
	port, err := strconv.ParseUint(os.Getenv(portEnv), 10, 32)
	if err != nil {
		return conf, err
	}
	conf.port = port
	conf.CertPath = os.Getenv(certPathEnv)
	conf.KeyPath = os.Getenv(keyPathEnv)
	return conf, nil
}
