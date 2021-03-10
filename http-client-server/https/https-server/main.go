package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

const (
	addr     = ":8443"
	certFile = "../etc/certs/test.crt"
	keyFile  = "../etc/certs/test.key"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := io.WriteString(writer, time.Now().String())
		if err != nil {
			log.Println(err)
			_, _ = io.WriteString(writer, err.Error())
		}
	})
	log.Fatal(http.ListenAndServeTLS(addr, certFile, keyFile, nil))
}
