package tests

import (
	"fmt"
	"testing"
)

const (
	A = iota
	B = iota
	C = iota
)

const (
	D, E, F = iota, iota, iota
)

const (
	X = iota
	Y
	_
	Z
)

func TestIota(t *testing.T) {
	fmt.Println(A, B, C)
	fmt.Println(D, E, F)
	fmt.Println(X, Y, Z)
}
