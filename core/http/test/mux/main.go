package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/x/", func(writer http.ResponseWriter, request *http.Request) {

		writer.Write([]byte(request.URL.String()))
	})

	http.ListenAndServe(":80", mux)
}
