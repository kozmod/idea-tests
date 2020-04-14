package datastruct

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func tryInitSlice(s []int) {
	s = make([]int, 0, 5)
	fmt.Println(fmt.Sprintf("m == nil in tryInitSlice?: %t, %v, len:%d, cap:%d", s == nil, s, len(s), cap(s)))
}

func TestSliceInit(t *testing.T) {
	var s []int
	tryInitSlice(s)
	fmt.Println(fmt.Sprintf("m == nil in TestSliceInit?: %t, %v, len:%d, cap:%d", s == nil, s, len(s), cap(s)))
	assert.Nil(t, s)
	s = append(s, 1, 2, 3)
	fmt.Println(fmt.Sprintf("m == nil after append?: %t, %v, len:%d, cap:%d", s == nil, s, len(s), cap(s)))
	assert.NotNil(t, s)
}

func TestChangeStructInSlice(t *testing.T) {
	a, b, x := "a", "b", "X!"

	s := []someStruct{
		2: {Val: a},
		0: {Val: b},
	}
	fmt.Println(s)
	assert.Equal(t, a, s[2].Val)
	s[2].Val = x //TODO not allowed in MAP, but in slice is allowed
	fmt.Println(s)
	assert.Equal(t, x, s[2].Val)
}
