package _struct

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type explodes interface {
	Bang() string
	Boom() string
}

// Type Bomb implements explodes
type Bomb struct{}

func (*Bomb) Bang() string {
	return "Bang"
}
func (Bomb) Boom() string {
	return "Boom"
}

func TestNilInterface(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			assert.NotNil(t, p)
			fmt.Println(p)
		}
	}()

	var bomb *Bomb = nil
	var explodes explodes = bomb
	println(bomb, explodes) // '0x0 (0x10a7060,0x0)'
	if explodes != nil {
		println(explodes.Bang()) // works fine
		println(explodes.Boom()) // panic: value method main.Bomb.Boom called using nil *Bomb pointer
	}
}
