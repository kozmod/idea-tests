package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var port string
	flag.StringVar(&port, "port", ":8080", "Server port")

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	httpMux := http.NewServeMux()
	httpMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprintf(writer, "time: %v\n", time.Now())
		if err != nil {
			log.Printf("error occured: %v", err)
		}
	})
	httpMux.HandleFunc("/stop", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "Try server\n")
		if err != nil {
			log.Printf("error occured: %v", err)
		}
		done <- syscall.SIGINT
	})

	httpServer := http.Server{
		Addr:    port,
		Handler: httpMux,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Printf("Server Started on port %s", port)

	<-done
	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")
}
