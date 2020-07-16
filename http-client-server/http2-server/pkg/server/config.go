package server

import (
	"github.com/spf13/viper"
)

type LogMod string

const (
	Stdout LogMod = "stdout"
	File          = "file"
)

type Config struct {
	ServerAddr struct {
		Port string
	}
	Log struct {
		Mod      []LogMod
		FilePath string
	}
}

func (c *Config) ColonPort() string {
	return ":" + c.ServerAddr.Port
}

func ReadConfig(cName, cType string, paths ...string) (c Config, err error) {
	viper.SetConfigName(cName)
	viper.SetConfigType(cType)
	for _, path := range paths {
		viper.AddConfigPath(path)
	}
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	if err = viper.Unmarshal(&c); err != nil {
		return
	}
	return
}
