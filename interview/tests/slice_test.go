package tests

import (
	"fmt"
	"testing"
)

func TestChaneSliceTest(t *testing.T) {
	a := []string{"a", "b", "c"}
	b := a[1:2]
	b[0] = "Foo"
	fmt.Println(a)

	c := []string{"a", "b", "c"}
	d := c[1:2]
	c = append(c, "x")
	d[0] = "Foo"
	fmt.Println(c)
}

func TestNilSlice(t *testing.T) {
	var s []string
	fmt.Println(s, len(s), cap(s))
	fmt.Println([]string(nil))

	s2 := append([]string(nil), []string(nil)...)
	fmt.Println(s2, len(s2), cap(s))
}
