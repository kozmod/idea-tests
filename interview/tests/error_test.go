package tests

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

type MyError struct{}

func (MyError) Error() string { return "MyError!" }

func errorHandler(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func TestPointerError(t *testing.T) {
	var err *MyError
	errorHandler(err)

	err = &MyError{}
	errorHandler(err)
}

func TestWrapAndCause(t *testing.T) {
	err := sql.ErrNoRows
	err = errors.Wrap(err, "wrapped")
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Printf("is sql.ErrNoRows: %+v\n", err)
	}
	fmt.Printf("cause: %+v\n", errors.Cause(err))
}
