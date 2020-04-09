package errors

import (
	"errors"
	"fmt"
	"testing"
)

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

var qe = &errorString{"A"}
var etp *errorString

func TestSwitchError(t *testing.T) {
	err := &errorString{"A"}
	switch {
	case err == nil:
		fmt.Println("NIL")
	case errors.Is(err, qe):
		fmt.Println("IS")
	case errors.As(err, &etp):
		fmt.Println("AS")
	case err == qe:
		fmt.Println("==")
	default:
		fmt.Println("DEFAULT")
	}
}
