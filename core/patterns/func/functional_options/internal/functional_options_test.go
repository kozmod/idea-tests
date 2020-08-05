package internal

import (
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
)

const addr = ":8080"

func TestFO_1(t *testing.T) {
	server := NewServer(addr)
	assert.Equal(t, server.addr, addr)
	assert.Equal(t, server.timeout, 0*time.Second)
	assert.Equal(t, len(server.whitelistIPs), 0)
}

func TestFO_2(t *testing.T) {
	server := NewServer(addr, Timeout(3*time.Second), WithWhitelistedIP("1"))
	assert.Equal(t, server.addr, addr)
	assert.Equal(t, server.timeout, 3*time.Second)
	assert.Equal(t, len(server.whitelistIPs), 1)
}
