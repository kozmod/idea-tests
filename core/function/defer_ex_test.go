package function

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"strconv"
	"testing"
)

func TestExDefer1(t *testing.T) {
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

func TestExDefer2(t *testing.T) {
	var i int
	i++
	defer assert.Equal(t, i, 1)
	defer ex(t, &i, 5)
	defer ex(t, &i, 4)
	defer func() {
		i++
		assert.Equal(t, i, 3)
	}()
	i++
	assert.Equal(t, i, 2)

}

func ex(t *testing.T, i *int, res int) {
	*i++
	assert.Equal(t, *i, res)
}
