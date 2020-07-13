package cmd

import (
	"github.com/kozmod/idea-tests/http-client-server/http2-server/pkg/server"
)

var (
	DefaultServerPort  = "8080"
	DefaultLogMod      = [...]server.LogMod{server.Stdout}
	DefaultLogFilePath = "http2server.log"
)
