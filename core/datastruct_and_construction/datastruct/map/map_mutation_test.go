package _map_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	one = 1
	two = 2
	ten = 10
)

func TestMapMutation(t *testing.T) {
	changeMap := func(m map[int]int) {
		m[ten] = two
	}

	m := make(map[int]int)
	m[ten] = one
	println("m[10] before change =", m[ten])
	assert.Equal(t, one, m[ten])
	changeMap(m)
	println("m[10] after change =", m[ten])
	assert.Equal(t, two, m[ten])
}

//noinspection ALL
func TestMapInitByFunc(t *testing.T) {
	tryInitMap := func(m map[int]int) {
		m = make(map[int]int)
		fmt.Println(fmt.Sprintf("m == nil in tryInitMap?: %t, %v", m == nil, m))
	}

	var m map[int]int
	tryInitMap(m)
	fmt.Println(fmt.Sprintf("m == nil in TestMapInit?: %t, %v", m == nil, m))
	assert.Nil(t, m)
}

func TestInit(t *testing.T) {
	m1 := make(map[int]int)
	m2 := new(map[int]int)
	var m3 map[int]int
	m4 := map[int]int{1: 1, 2: 1}
	assert.NotNil(t, m1)
	assert.NotNil(t, &m1)
	assert.NotNil(t, m2)
	assert.NotNil(t, &m2)
	assert.Nil(t, *m2)
	assert.Nil(t, m3)
	assert.NotNil(t, &m3)
	assert.NotNil(t, m4)
	assert.NotNil(t, &m4)
	fmt.Println(fmt.Sprintf("m1: %v", m1))
	fmt.Println(fmt.Sprintf("m2: %v", m2))
	fmt.Println(fmt.Sprintf("m3: %v", m3))
	fmt.Println(fmt.Sprintf("m4: %v", m4))
}

func TestChangeStructInMap(t *testing.T) {
	type SomeStruct struct {
		Val string
	}
	m := map[int]SomeStruct{
		1: {Val: "aaa"},
	}
	/***********************
		not allowed in MAP
	************************/
	//m[1].Val ="bbb"
	assert.NotNil(t, m)
}
