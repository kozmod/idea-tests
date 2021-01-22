package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	logFile    = "./" + time.Now().Format("02_01_2006__15_04_05") + ".log"
	serverPort = ":8099"
)

func init() {
	//log from ENV
	logFileEnv := os.Getenv("LOG_FILE")
	logFile = strings.ReplaceAll(logFileEnv, " ", "")
	switch {
	case strings.ToUpper(logFileEnv) == "OFF":
		log.SetOutput(ioutil.Discard)
	case strings.ToUpper(logFileEnv) != "STDOUT":
		if logFileEnv != "" {
			logFile = logFileEnv
		}
		if f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); err != nil {
			log.Fatalf("error opening file: %v", err)
		} else {
			log.SetOutput(io.MultiWriter(os.Stdout, f))
		}
	}

	//port from ENV
	portEnv := os.Getenv("SERVER_PORT")
	if portEnv = strings.ReplaceAll(portEnv, " ", ""); portEnv != "" {
		if _, err := strconv.Atoi(portEnv); err != nil {
			panic(fmt.Sprintf("port not valid: %v", err))
		} else {
			serverPort = ":" + portEnv
		}
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("incoming rq:%v", r)
		proxyReq, err := http.NewRequest(r.Method, r.URL.String(), r.Body)
		if err != nil {
			log.Fatalf("create proxyReq error: %v", err)
		}

		CopyHeaders(proxyReq.Header, r.Header)

		client := &http.Client{}
		proxyRes, err := client.Do(proxyReq)
		if err != nil {
			log.Fatalf("err proxy request: %v", err)
		}
		resBody := proxyRes.Body
		defer func() {
			if err := resBody.Close(); err != nil {
				log.Fatalf("close proxyResponse error:%v\n", err)
			}
		}()

		CopyHeaders(w.Header(), proxyRes.Header)

		if _, err := io.Copy(w, resBody); err != nil {
			log.Fatalf("copy proxyResponse error:%v\n", err)
		}
	})
	err := http.ListenAndServe(serverPort, nil)
	if err != nil {
		log.Fatalf("ListenAndServe error: %v", err)
	}
}

func CopyHeaders(src, dst http.Header) {
	for header, values := range src {
		for _, value := range values {
			dst.Add(header, value)
		}
	}
}
