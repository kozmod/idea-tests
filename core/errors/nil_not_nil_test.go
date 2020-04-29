package errors

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type customError struct{}

// Error implements the error interface.
func (c *customError) Error() string {
	return "Find the bug."
}

// fail returns nil values for both return types.
func fail() ([]byte, *customError) {
	return nil, nil
}

func fail2() ([]byte, error) {
	return nil, nil
}

// Why is my nil error value not equal to nil?
//https://golang.org/doc/faq#nil_error
//https://stackoverflow.com/questions/31728656/why-is-err-nil
//https://stackoverflow.com/questions/58272987/assigning-nil-value-to-driver-value-makes-driver-value-not-equal-nil
//https://stackoverflow.com/questions/53892508/golang-returning-nil-does-not-return-nil
//https://stackoverflow.com/questions/44900065/nil-pointer-to-struct-not-deep-equal-to-nil
func TestNilOrNotNil(t *testing.T) {
	var err error
	if _, err = fail(); err != nil {
		assert.Nil(t, err)
		log.Println("Why did this fail?")
	} else {
		log.Println("`fail()` err is nil")
	}

	if _, err = fail2(); err != nil {
		log.Println("Why did this fail?")
	} else {
		log.Println("`fail2()` err is nil")
	}
}
