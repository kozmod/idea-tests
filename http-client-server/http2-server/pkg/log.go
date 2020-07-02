package pkg

import (
	"github.com/kozmod/idea-tests/http-client-server/http2-server/pkg/server"

	"io"
	"log"
	"os"

	. "github.com/kozmod/idea-tests/http-client-server/http2-server/cmd"
)

func init() {
	config := Config
	var writers []io.Writer
	//if Contains(DefaultLogMod[:], File) {
	if Contains(config.Log.Mod[:], server.File) {
		// log to console and file
		f, err := os.OpenFile(config.Log.FilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		writers = append(writers, f)
	}
	if Contains(config.Log.Mod[:], server.Stdout) {
		writers = append(writers, os.Stdout)
	}
	if len(writers) < 1 {
		log.Fatal("logger have to set 1 writer at least")
	}
	log.SetOutput(io.MultiWriter(writers...))
}

func Contains(mods []server.LogMod, search server.LogMod) bool {
	for _, mod := range mods {
		if mod == search {
			return true
		}
	}
	return false
}
