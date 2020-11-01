package main

import (
	"net/http"
	"net/http/pprof"
	//_ "net/http/pprof"
)

func main() {
	//http://localhost:9090/debug/pprof/
	//http.ListenAndServe(":9090", nil)

	pprofMux := http.NewServeMux()
	pprofMux.HandleFunc("/debug/pprof/", pprof.Index)
	pprofMux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	pprofMux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	pprofMux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	pprofMux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	http.ListenAndServe(":9090", pprofMux)
}
