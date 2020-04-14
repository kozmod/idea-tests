package pack

import "fmt"

var Val string

func init() {
	fmt.Println("INIT ROOT")
	Val = "sub"
}

func getVal() string {
	fmt.Println("GET VAL - ROOT")
	return Val
}
