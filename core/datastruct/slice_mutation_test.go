package datastruct

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func tryInitSlice(s []int) {
	s = make([]int, 0, 5)
	fmt.Println(fmt.Sprintf("s == nil in tryInitSlice?: %t, %v, len:%d, cap:%d", s == nil, s, len(s), cap(s)))
}

func TestSliceInit(t *testing.T) {
	var s []int
	tryInitSlice(s)
	fmt.Println(fmt.Sprintf("s == nil in TestSliceInit?: %t, %v, len:%d, cap:%d", s == nil, s, len(s), cap(s)))
	assert.Nil(t, s)
	s = append(s, 1, 2, 3)
	fmt.Println(fmt.Sprintf("s == nil after append?: %t, %v, len:%d, cap:%d", s == nil, s, len(s), cap(s)))
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

	/***********************
		not allowed in MAP, but in SLICE is allowed
	************************/
	s[2].Val = x
	fmt.Println(s)
	assert.Equal(t, x, s[2].Val)
}

func appendToSlice(s []int, i int) []int {
	s = append(s, i)
	fmt.Println("***************************************************")
	fmt.Println(fmt.Sprintf("s   in appendToSlice?:%v, len:%d, cap:%d", s, len(s), cap(s)))
	fmt.Println("***************************************************")
	return s
}

func TestAppendToSliceInFunc(t *testing.T) {
	s := make([]int, 0, 5)
	rs := appendToSlice(s, 1)

	fmt.Println(fmt.Sprintf("s   in TestAppendToSliceInFunc?:%v, len:%d, cap:%d", s, len(s), cap(s)))
	fmt.Println(fmt.Sprintf("rs  in TestAppendToSliceInFunc?:%v, len:%d, cap:%d", rs, len(rs), cap(rs)))

	s2 := make([]int, 0, 5)
	rs2 := append(s2, 2)
	fmt.Println(fmt.Sprintf("s2  in TestAppendToSliceInFunc?:%v, len:%d, cap:%d", s2, len(s2), cap(s2)))
	fmt.Println(fmt.Sprintf("rs2 in TestAppendToSliceInFunc?:%v, len:%d, cap:%d", rs2, len(rs2), cap(rs2)))
}
