package viper

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	CollectorClientConfig
	ColocatorServerConfig
}

type CollectorClientConfig struct {
	GrpcHost string
	GrpcPort string
	T        time.Duration
}

type ColocatorServerConfig struct {
	GrpsPort int
}

func TestReadConfig(t *testing.T) {
	var Config Config
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(filepath.Join("../viper"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	spew.Dump(Config)
	Config.T = Config.T * time.Minute
	spew.Dump(Config)
	fmt.Println(Config)
}

type ExecuteJidResp struct {
	Info   []interface{}
	Return []map[string]interface{}
}

var yaml = `info: []
return:
  - foo.com:
      some_value:
        1.1.1.1: true
      some_value2:
        hey: hi`

func TestReadConfigFromString(t *testing.T) {
	var Resp ExecuteJidResp
	viper.SetConfigType("yml")

	if err := viper.ReadConfig(strings.NewReader(yaml)); err != nil {
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&Resp); err != nil {
		log.Fatalln(err)
	}
	spew.Dump(Resp)
	fmt.Println(Resp)
}
