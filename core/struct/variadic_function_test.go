package _struct

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	apply(nil)

}

func apply(arg ...string) {
	for _, o := range arg {
		fmt.Println(o)
	}
}
