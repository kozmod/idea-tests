package sub

import "fmt"

var Val string

func init() {
	fmt.Println("INIT SUB")
	Val = "sub"
}

func getVal() string {
	fmt.Println("GET VAL - SUB")
	return Val
}
