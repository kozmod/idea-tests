package _type

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestNegativeIntToUint(t *testing.T) {
	var i int = -1
	fmt.Println(uint(i))
	assert.Equal(t, uint(i), uint(18446744073709551615))
}
