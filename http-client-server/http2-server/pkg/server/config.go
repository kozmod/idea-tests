package server

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
)

type Config struct {
	ServerAddr struct {
		Port string
	}
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
	spew.Dump(c)
	return
}
