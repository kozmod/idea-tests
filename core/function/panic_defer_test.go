package function

import (
	"errors"
	"fmt"
	"runtime/debug"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPanicDefer1(t *testing.T) {
	defer debug.PrintStack()
	f(3)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // Сбой при х = 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

func TestPanicDefer_Division0(t *testing.T) {
	res, err := division(2, 2)
	assert.Equal(t, 1, res)
	assert.Nil(t, err)

	res, err = division(2, 0)
	assert.Equal(t, 0, res)
	assert.NotNil(t, err)
}

func division(x, y int) (res int, err error) {
	defer func() {
		if p := recover(); p != nil {
			err = errors.New("деление на 0")
		}
	}()
	res = x / y
	return
}

func TestPanicDefer_AbsPanic(t *testing.T) {
	res := absPanic(1)
	assert.Equal(t, uint(1), res)

	res = absPanic(0)
	assert.Equal(t, uint(0), res)

	res = absPanic(-1)
	assert.Equal(t, uint(1), res)

}

func absPanic(x int) (res uint) {
	defer func() {
		if p := recover(); p != nil {
			tmp, _ := p.(int)
			res = uint(tmp * -1)
		}
	}()
	if x < 0 {
		panic(x)
	}
	return uint(x)
}
