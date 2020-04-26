package pkg

import (
	"github.com/kozmod/idea-tests/http-client-server/http2-client/cmd"
	"io"
	"log"
	"os"
)

func init() {
	// log to console and file
	f, err := os.OpenFile(cmd.DefaultLogFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	wrt := io.MultiWriter(os.Stdout, f)

	log.SetOutput(wrt)
}