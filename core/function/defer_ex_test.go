package function

import (
	"fmt"
	"strconv"
	"testing"
)

func TestExDefer(t *testing.T) {
	{
		defer func() {
			fmt.Println("from {} defer")
		}() // не будет закрыто в конце этого блока
	}
	for i := 0; i < 2; i++ {
		defer func(v int) {
			fmt.Println("for fori defer: " + strconv.Itoa(v))
		}(i) // не будет закрыто в конце этого блока
	}
	fmt.Println("from Test")
	defer func() {
		fmt.Println("From Test defer")
	}()
}
