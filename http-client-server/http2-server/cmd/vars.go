package cmd

import (
	"log"

	"github.com/kozmod/idea-tests/http-client-server/http2-server/pkg/server"
)

var (
	Config server.Config

	DefaultServerPort  = "8080"
	DefaultLogMod      = [...]server.LogMod{server.Stdout}
	DefaultLogFilePath = "http2server.log"
)

func init() {
	if conf, err := server.ReadConfig("config", "yml", "./etc/config"); err == nil {
		Config = conf
	} else {
		log.Println(err)
		log.Println("get default config from \"github.com/kozmod/idea-tests/http-client-server/http2-server/cmd\"")
		Config = server.Config{
			ServerAddr: struct{ Port string }{Port: DefaultServerPort},
			Log: struct {
				Mod      []server.LogMod
				FilePath string
			}{Mod: DefaultLogMod[:],
				FilePath: DefaultLogFilePath,
			},
		}
	}
}
