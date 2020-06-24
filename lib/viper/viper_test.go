package viper

import (
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
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
	//fmt.Println(Config)
}
