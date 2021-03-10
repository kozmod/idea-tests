package _struct

import (
	"fmt"
	"testing"
)

type I interface {
	do()
}

type P struct {
	val string
}

func (x *P) do() {
}

type C struct {
	val string
}

func (x C) do() {
}

func do(s I) {
	fmt.Println(s)
	fmt.Println(&s)
}

func changePointerValue(x *P) {
	*x = P{"XXXX"}
}

func TestInterfaceAsparam(t *testing.T) {
	fmt.Println("P{\"p\"} : CHANGE VALUE : ----------->")
	p0 := P{"p"}
	changePointerValue(&p0)
	fmt.Println(p0)

	fmt.Println("P{\"p\"} : ----------->")
	p := P{"p"}
	do(&p)

	fmt.Println("C{\"c\"} : ----------->")
	c := C{"c"}
	do(c)
	do(&c)

	fmt.Println("&C{\"c\"} : ----------->")
	c2 := &C{"c"}
	do(c2)
	do(*c2)
	//do(&c2)//compile error
}
