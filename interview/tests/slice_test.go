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
