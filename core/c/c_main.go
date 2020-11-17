package main

/*
int Multiply(int a, int b) { // need to gcc (Linux)
	return a * b;
}
*/
import "C"
import "errors"

func main() {
	a := 2
	b := 3
	res := C.Multiply(C.int(a), C.int(b))
	if a*b != int(res) { //need to cast to go int
		panic(errors.New("C problem"))
	}
}
