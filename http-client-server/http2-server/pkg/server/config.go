package server

import (
	"log"
	"path/filepath"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
)

type Config struct {
	ServerAddr struct {
		Port string
	}
}

func ReadConfig(cname, ctype string, paths ...string) (c Config, ok bool) {
	viper.SetConfigName(cname)
	viper.SetConfigType(ctype)
	viper.AddConfigPath(filepath.Join(paths...))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
		ok = false
		return
	}

	if err := viper.Unmarshal(&c); err != nil {
		log.Println(err)
		ok = false
		return
	}
	spew.Dump(c)
	return
}
