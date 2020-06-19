package server

import (
	"fmt"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func ConfigureAndServe(port string) {
	ConfigureHandleFuncsAndServe(
		port,
		map[string]func(http.ResponseWriter, *http.Request){
			"/t": currentTimePayload,
			"/h": headersPayload,
			"/":  helloCurrentTime,
		},
	)
}

func ConfigureHandleFuncsAndServe(port string, handleMap map[string]func(http.ResponseWriter, *http.Request)) {
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
}

func currentTimePayload(w http.ResponseWriter, req *http.Request) {
	content := logAndGetContent(w, req)
	fmt.Fprintf(w, "time: %v\npayload:%s\n", time.Now(), content)
}

func helloCurrentTime(w http.ResponseWriter, req *http.Request) {
	logAndGetContent(w, req)
	fmt.Fprintf(w, "Hellow,\nCurrent time: %v\n", time.Now())
}

func headersPayload(w http.ResponseWriter, req *http.Request) {
	content := logAndGetContent(w, req)
	fmt.Fprintf(w, "/*********************************\n")
	fmt.Fprintf(w, "|---------|\n")
	fmt.Fprintf(w, "| Headers |\n")
	fmt.Fprintf(w, "|---------|\n")
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
	fmt.Fprintf(w, "*********************************/\n")
	fmt.Fprintf(w, "|---------|\n")
	fmt.Fprintf(w, "| Payload |\n")
	fmt.Fprintf(w, "|---------|\n")
	fmt.Fprintf(w, "%s", content)
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
