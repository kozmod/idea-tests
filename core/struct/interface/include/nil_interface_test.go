package include

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const nilVal = "0"
const notNilVal = "hello"

type i interface {
	do()
	val() string
}

type st struct {
	str string
}

func (s *st) do() {
	if s == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(s.str)
}

func (s *st) val() (val string) {
	if s == nil {
		return nilVal
	}
	return s.str
}

func TestNilInterface_1(t *testing.T) {
	var i i
	var s *st
	assert.Nil(t, s)
	assert.Nil(t, i)

	i = s
	assert.Nil(t, s)
	assert.Nil(t, i)

	describe(i)
	i.do()
	assert.Equal(t, nilVal, i.val())

	i = &st{notNilVal}
	describe(i)
	i.do()
	assert.Equal(t, notNilVal, i.val())
}

func describe(i i) {
	fmt.Printf("(%v, %T)\n", i, i)
}
