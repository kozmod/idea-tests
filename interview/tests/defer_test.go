package tests

import (
	"fmt"
	"testing"
)

func TestDeferInit(t *testing.T) {
	x := 5
	defer func() {
		fmt.Println(x)
	}()
	for i := 0; i < 10; i++ {
		x++
	}
}

//goland:noinspection GoUnreachableCode
func TestDeferPanic(t *testing.T) {
	panic("some panic")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
}
