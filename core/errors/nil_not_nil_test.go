/*
Why is my nil error value not equal to nil?

	https://golang.org/doc/faq#nil_error
	https://stackoverflow.com/questions/31728656/why-is-err-nil
	https://stackoverflow.com/questions/58272987/assigning-nil-value-to-driver-value-makes-driver-value-not-equal-nil
	https://stackoverflow.com/questions/53892508/golang-returning-nil-does-not-return-nil
	https://stackoverflow.com/questions/44900065/nil-pointer-to-struct-not-deep-equal-to-nil
*/
package errors

import (
	"fmt"
	"testing"
)

type customError struct{}

// Error implements the error interface.
func (*customError) Error() string {
	return "Find the bug."
}

// fail returns nil values for both return types.
func failA() ([]byte, *customError) {
	return nil, nil
}

func failB() ([]byte, error) {
	return nil, nil
}

func TestNilOrNotNil(t *testing.T) {
	var err error

	step := 1
	if _, err = failA(); err != nil {
		fmt.Printf("%d:`failA()` -> Why did this fail?: %v\n", step, err)
	} else {
		fmt.Printf("%d:`failA()` -> err: %v\n", step, err)
	}

	step = 2
	if _, err := failA(); err != nil { // lock at
		fmt.Printf("%d:`failA()` -> Why did this fail?: %v\n", step, err)
	} else {
		fmt.Printf("%d:`failA()` -> err: %v\n", step, err)
	}

	step = 3
	if _, err = failB(); err != nil {
		fmt.Printf("%d:`failB()` ->Why did this fail?: %v\n", step, err)
	} else {
		fmt.Printf("%d:`failB()` -> err: %v\n", step, err)
	}

	step = 4
	if _, err := failB(); err != nil {
		fmt.Printf("%d:`failB()` ->Why did this fail?: %v\n", step, err)
	} else {
		fmt.Printf("%d:`failB()` -> err: %v\n", step, err)
	}
}
