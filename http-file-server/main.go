package main

import (
	"flag"
	"log"
	"net/http"
)

var port = flag.String("port", ":8181", "Port of the server")
var dir = flag.String("dir", "", "Dir of the server (may be absolute or relative)")

func main() {
	flag.Parse()
	log.Fatal(http.ListenAndServe(*port, http.FileServer(http.Dir(*dir))))
}
