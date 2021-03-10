package config_test

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
)

func TestGetConfigEnv(t *testing.T) {
	exSA := "1"
	exRQ := "2"
	exRF := "3"
	exPPU := "4"

	_ = os.Setenv(ServerAddrEnv, exSA)
	_ = os.Setenv(RequestQuantityEnv, exRQ)
	_ = os.Setenv(RequestFrequencySecEnv, exRF)
	_ = os.Setenv(PostWithPayloadUrlEnv, exPPU)

	conf := FromEnv()
	assert.Equal(t, exSA, conf.ServerAddr())
	q, _ := strconv.Atoi(exRQ)
	assert.Equal(t, q, conf.RequestQuantity())
	assert.Equal(t, asSeconds(exRF), conf.RequestFrequency())
	assert.Equal(t, exPPU, conf.PostWithPayloadUrl())
}

func asSeconds(sec string) time.Duration {
	f, _ := strconv.Atoi(sec)
	return time.Duration(f) * time.Second
}
