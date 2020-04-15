package _struct

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type dataEq struct {
	num     int
	fp      float32
	complex complex64
	str     string
	char    rune
	yes     bool
	events  <-chan string
	handler interface{}
	ref     *byte
	raw     [10]byte
}

type data struct {
	num    int               // ok
	checks [10]func() bool   // несравниваемо
	doit   func() bool       // несравниваемо
	m      map[string]string // несравниваемо
	bytes  []byte            // несравниваемо
}

func TestEquals(t *testing.T) {
	v1 := dataEq{}
	v2 := dataEq{}
	assert.True(t, v1 == v2) // выводит: v1 == v2: true
}

func TestNotEquals(t *testing.T) {
	v1 := data{}
	v2 := data{}
	_ = v1
	_ = v2
	// fmt.Println("v1 == v2:", v1 == v2)// invalid operation: v1 == v2 (struct containing [10]func() bool cannot be compared)
}
