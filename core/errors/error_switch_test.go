package errors

import (
	"fmt"
	"github.com/pkg/errors"
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

func TestWrap(t *testing.T) {
	err := &errorString{"A"}
	err2 := errors.Wrap(err, "Mongo")

	fmt.Println(err2)
	fmt.Println(errors.Unwrap(err2))
}
