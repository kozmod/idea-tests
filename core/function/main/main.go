package main

import (
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	http.ListenAndServe(":8080", nil)
	var slice []string
	i := 0
	for {
		i++
		time.Sleep(1 * time.Microsecond)
		slice = append(slice, string(rune(rand.Int())))
		if i == 50 {
			slice = make([]string, 0)
		}

	}
}
