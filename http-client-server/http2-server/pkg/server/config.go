package server

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	ServerAddr struct {
		Port string
	}
}

func ReadConfig(cname, ctype string, paths ...string) (c Config, ok bool) {
	viper.SetConfigName(cname)
	viper.SetConfigType(ctype)
	for _, path := range paths {
		viper.AddConfigPath(path)
	}
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
