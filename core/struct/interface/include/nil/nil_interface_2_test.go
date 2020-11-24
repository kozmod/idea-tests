package nil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	bang = "Bang"
	boom = "Boom"
)

type explodes interface {
	Bang() string
	Boom() string
}

// Type Bomb implements explodes
type Bomb struct{}

func (*Bomb) Bang() string {
	return bang
}
func (Bomb) Boom() string {
	return boom
}

func TestNilInterface_2(t *testing.T) {
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

type explodesSeries struct {
	bomb        Bomb
	bombPtr     *Bomb
	explodes    explodes
	explodesPrt *explodes
}

func TestNilInterface_3(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			assert.NotNil(t, p)
		}
	}()

	explodes := explodesSeries{}

	assert.NotNil(t, explodes.bomb)
	assert.Nil(t, explodes.bombPtr)
	assert.Nil(t, explodes.explodes)
	assert.Nil(t, explodes.explodesPrt)

	assert.Equal(t, explodes.bomb.Bang(), bang)
	assert.Equal(t, explodes.bomb.Boom(), boom)

	assert.Equal(t, explodes.bombPtr.Bang(), bang)
	assert.Equal(t, explodes.bombPtr.Boom(), boom) // panic: value method main.Bomb.Boom called using nil *Bomb pointer

	assert.Equal(t, explodes.explodes.Bang(), bang) // nil -> invalid memory address or nil pointer dereference
	assert.Equal(t, explodes.explodes.Boom(), boom) // nil -> invalid memory address or nil pointer dereference

	//assert.Equal(t,explodes.explodesPrt.Bang(), bang) // Unresolved reference 'Bang'
	//assert.Equal(t,explodes.explodesPrt.Boom(), boom) // Unresolved reference 'Bang'
}
