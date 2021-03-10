package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/kozmod/idea-tests/http-client-server/http2-client/pkg/utils"
)

const (
	ServerAddrEnv          = "SERVER_ADDR"
	RequestQuantityEnv     = "REQUEST_QUANTITY"
	RequestFrequencySecEnv = "REQUEST_FREQUENCY_SEC"
	PostWithPayloadUrlEnv  = "POST_WITH_PAYLOAD_URL"
)

type Config struct {
	serverAddr         string
	requestQuantity    int
	requestFrequency   time.Duration
	postWithPayloadUrl string
}

func (c *Config) ServerAddr() string {
	return c.serverAddr
}

func (c *Config) RequestQuantity() int {
	return c.requestQuantity
}

func (c *Config) RequestFrequency() time.Duration {
	return c.requestFrequency
}

func (c *Config) PostWithPayloadUrl() string {
	return c.postWithPayloadUrl
}

func FromEnv() Config {
	conf := Config{}
	conf.serverAddr = os.Getenv(ServerAddrEnv)
	q, err := strconv.Atoi(os.Getenv(RequestQuantityEnv))
	if err != nil {
		log.Fatal(err)
	}
	conf.requestQuantity = q
	f, err := utils.AsSeconds(os.Getenv(RequestFrequencySecEnv))
	if err != nil {
		log.Fatal(err)
	}
	conf.requestFrequency = f
	conf.postWithPayloadUrl = os.Getenv(PostWithPayloadUrlEnv)
	return conf
}
