package errors

import (
	"database/sql"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/pkg/errors"
)

type errorString struct {
	msg string
	val string
}

func (e *errorString) Error() string {
	return e.msg
}

func newErrorString() error {
	return &errorString{msg, st}
}

type errorOne struct{}

func (e errorOne) Error() string {
	return "Error One"
}

const (
	msg = "MSG"
	st  = "SOME_TEXT"
)

var (
	qe = newErrorString()
	se = errors.New("some error")
)

//goland:noinspection ALL
func TestIs(t *testing.T) {
	e := errors.New("a")
	e2 := errors.New("a")
	assert.False(t, errors.Is(e, e2))
	assert.False(t, errors.Is(e2, e))
	assert.False(t, e == e2)
	assert.False(t, &e == &e2)

	e3 := errorString{msg, st}
	e4 := errorString{msg, st}
	assert.False(t, errors.Is(&e3, &e4))
	assert.True(t, e3 == e4) //equals by value

	e5 := errorOne{}
	e6 := fmt.Errorf("E2: %w", errorOne{})
	assert.True(t, errors.Is(e6, e5))
	assert.False(t, errors.Is(e5, e6))
	assert.False(t, errors.Is(e6, &e5))

	e7 := se
	assert.True(t, errors.Is(e7, se))
	assert.True(t, errors.Is(se, e7))
	assert.False(t, errors.Is(e7, fmt.Errorf("E2: %w", se)))
	assert.True(t, errors.Is(fmt.Errorf("E2: %w", se), e7))
	assert.True(t, e7 == se)
}

func TestSwitchError(t *testing.T) {
	err := newErrorString()
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

func TestAs(t *testing.T) {
	err := newErrorString()
	var es *errorString
	if errors.As(err, &es) {
		fmt.Printf("Err[msg: %s; val: %s]\n", es.msg, es.val)
		assert.Equal(t, msg, es.msg)
		assert.Equal(t, st, es.val)
	}
}

func TestEquality(t *testing.T) {
	sqle := sql.ErrNoRows
	assert.True(t, sqle == sql.ErrNoRows)
	assert.False(t, sqle == fmt.Errorf("foo err, %v", sql.ErrNoRows))
	assert.True(t, sql.ErrNoRows == errors.Cause(sqle))
}
