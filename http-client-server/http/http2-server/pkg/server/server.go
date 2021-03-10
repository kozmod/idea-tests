package server

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var HandleFunctionMap = map[string]func(http.ResponseWriter, *http.Request){
	"/hello": helloCurrentTime,
	"/ping":  ping,
	"/t":     currentTime,
	"/hp":    headersPayload,
}

func ConfigureAndServe(port string) http.Server {
	return ConfigureHandleFuncsAndServe(port, HandleFunctionMap)
}

func ConfigureHandleFuncsAndServe(port string, handleMap map[string]func(http.ResponseWriter, *http.Request)) http.Server {
	h2s := http2.Server{}
	hs := http.Server{
		Addr:    port,
		Handler: h2c.NewHandler(http.DefaultServeMux, &h2s),
	}
	_ = http2.ConfigureServer(&hs, &h2s)

	for k, v := range handleMap {
		http.HandleFunc(k, v)
	}

	log.Printf("Server run on %s", port)
	log.Fatal(hs.ListenAndServe())
	return hs
}

func currentTime(w http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprintf(w, "time: %v\n", time.Now())
}

func ping(w http.ResponseWriter, req *http.Request) {
	_ = logAndGetContent(w, req)
	_, _ = fmt.Fprintf(w, "pong")
}

func helloCurrentTime(w http.ResponseWriter, req *http.Request) {
	logAndGetContent(w, req)
	_, _ = io.WriteString(w, helloBox)
	_, _ = fmt.Fprintf(w, "Current time: %v\n", time.Now())
}

func headersPayload(w http.ResponseWriter, req *http.Request) {
	content := logAndGetContent(w, req)
	_, _ = io.WriteString(w, headerBox)
	for name, headers := range req.Header {
		for _, h := range headers {
			_, _ = fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
	_, _ = io.WriteString(w, payloadBox)
	_, _ = io.WriteString(w, content)
}

func logAndGetContent(w http.ResponseWriter, req *http.Request) string {
	log.Printf("Request connection: %s, path: %s, method: %s", req.Proto, req.URL.Path[1:], req.Method)
	defer req.Body.Close()
	if contents, err := ioutil.ReadAll(req.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalf("Oops! Failed reading body of the request.\n %s", err)
	} else {
		return string(contents)
	}
	return "fatal error"
}
