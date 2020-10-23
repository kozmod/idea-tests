package tests

import "testing"

func TestChangePointer(t *testing.T) {

	changePointer := func(p *int) {
		v := 3
		p = &v
	}

	v := 5
	p := &v
	println(*p)

	changePointer(p)
	println(*p)

}

func TestChangePointer_fix(t *testing.T) {
	changePointer := func(p *int) {
		v := 3
		*p = v
	}

	v := 5
	p := &v
	println(*p)

	changePointer(p)
	println(*p)
}
