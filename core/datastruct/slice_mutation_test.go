package datastruct

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

//noinspection GoNilness
func TestSliceInit(t *testing.T) {
	tryInitSlice := func(s []int) {
		s = make([]int, 0, 5)
		fmt.Println(fmt.Sprintf("s == nil in tryInitSlice?: %t, %v, len:%d, cap:%d", s == nil, s, len(s), cap(s)))
	}

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

func TestAppendToSliceInFunc(t *testing.T) {
	appendToSlice := func(s []int, i int) []int {
		s = append(s, i)
		fmt.Println("***************************************************")
		fmt.Println(fmt.Sprintf("s   in appendToSlice?:%v, len:%d, cap:%d", s, len(s), cap(s)))
		fmt.Println("***************************************************")
		return s
	}

	s := make([]int, 0, 5)
	rs := appendToSlice(s, 1)
	fmt.Println(fmt.Sprintf("s   in TestAppendToSliceInFunc?:%v, len:%d, cap:%d", s, len(s), cap(s)))
	fmt.Println(fmt.Sprintf("rs  in TestAppendToSliceInFunc?:%v, len:%d, cap:%d", rs, len(rs), cap(rs)))

	s2 := make([]int, 0, 0)
	rs2 := append(s2, 2)
	fmt.Println(fmt.Sprintf("s2  in TestAppendToSliceInFunc?:%v, len:%d, cap:%d", s2, len(s2), cap(s2)))
	fmt.Println(fmt.Sprintf("rs2 in TestAppendToSliceInFunc?:%v, len:%d, cap:%d", rs2, len(rs2), cap(rs2)))
}

func TestAppendToSliceInFunc_2(t *testing.T) {
	newSliceAndAppend := func(value []string) {
		fmt.Printf("value=%v\n", value)

		value2 := value[:]
		value2 = append(value2, "b")
		fmt.Printf("value=%v, value2=%v\n", value, value2)

		value2[0] = "z"
		fmt.Printf("value=%v, value2=%v\n", value, value2)
	}

	slice1 := []string{"a"} // length 1, capacity 1

	newSliceAndAppend(slice1)
	/*
		Output:
		value=[a] -- ok
		value=[a], value2=[a b] -- ok: value unchanged, value2 updated
		value=[a], value2=[z b] -- ok: value unchanged, value2 updated
	*/

	slice10 := make([]string, 1, 10) // length 1, capacity 10
	slice10[0] = "a"

	newSliceAndAppend(slice10)
	/*
		Output:
		value=[a] -- ok
		value=[a], value2=[a b] -- ok: value unchanged, value2 updated
		value=[z], value2=[z b] -- WTF?!? value changed???
	*/
}

func TestAppendToSliceInFunc_3(t *testing.T) {
	s1 := []int{1, 2, 3}
	fmt.Println(len(s1), cap(s1), s1) // выводит 3 3 [1 2 3]

	s2 := s1[1:]
	fmt.Println(len(s2), cap(s2), s2) // выводит 2 2 [2 3]

	for i := range s2 {
		s2[i] += 20
	}

	// всё ещё ссылается на тот же массив
	fmt.Println(s1) // выводит [1 22 23]
	fmt.Println(s2) // выводит [22 23]

	s2 = append(s2, 4) //append -> create new array to s2

	for i := range s2 {
		s2[i] += 10
	}

	//s1 is now "stale"
	fmt.Println(s1) // выводит [1 22 23]
	fmt.Println(s2) // выводит [32 33 14]
}
