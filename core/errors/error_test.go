package errors

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"github.com/pkg/errors"
	"testing"
)

type errorString struct {
	msg string
	val string
}

func (e *errorString) Error() string {
	return e.msg
}

func newError() error {
	return &errorString{msg, st}
}

const (
	msg = "MSG"
	st  = "SOME_TEXT"
)

var qe = newError()

func TestSwitchError(t *testing.T) {
	err := newError()
	var etp *errorString
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

func TestWrapAndUnwrap(t *testing.T) {
	err := newError()
	err2 := errors.Wrap(err, "Mongo")
	fmt.Println(err2)
	//fmt.Println(errors.Unwrap(err2))

	var etp *errorString
	if errors.As(err, &etp) {
		fmt.Printf("err is errorString: %s, %s \n", etp.msg, etp.val)
	} else {
		fmt.Println("err is not errorString")
	}
	if errors.As(err2, &etp) {
		fmt.Printf("err2 is errorString: %s, %s \n", etp.msg, etp.val)
	} else {
		fmt.Println("err2 is not errorString")
	}
}

func TestAs(t *testing.T) {
	err := newError()
	var es *errorString
	if errors.As(err, &es) {
		fmt.Printf("Err[msg: %s; val: %s]\n", es.msg, es.val)
		assert.Equal(t, msg, es.msg)
		assert.Equal(t, st, es.val)
	}
}
