package cmd

type LogMod string

const (
	Stdout LogMod = "stdout"
	File          = "file"
)

var (
	DefaultServerPort  = "8080"
	DefaultLogMod      = [...]LogMod{Stdout}
	DefaultLogFilePath = "http2server.log"
)
