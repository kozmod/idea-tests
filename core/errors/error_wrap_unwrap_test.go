package errors

import (
	"database/sql"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/pkg/errors"
)

type RepositoryError struct {
	errorMessage string
	cause        error
}

func (repErr *RepositoryError) Error() string {
	return repErr.errorMessage
}

func (repErr *RepositoryError) Cause() error {
	return repErr.cause
}

type OtherError struct {
	errorMessage string
	cause        error
}

func (repErr *OtherError) Error() string {
	return repErr.errorMessage
}

func (repErr *OtherError) Cause() error {
	return repErr.cause
}

func TestWrapAndUnwrap(t *testing.T) {
	repo := &RepositoryError{"MSG_R", nil}
	repo2 := errors.Wrap(repo, "Mongo")

	fmt.Println(repo2)
	var e *RepositoryError
	fmt.Println(errors.As(repo2, &e))

	other := &OtherError{"MSG_R", repo2}

	var oe *OtherError
	fmt.Println(errors.As(other, &oe))
	fmt.Println(errors.As(errors.Unwrap(other), &oe))
	fmt.Println(errors.Cause(other))
}

func TestWrap(t *testing.T) {
	we := errors.Wrap(sql.ErrNoRows, "foo failed")
	wme := errors.WithMessage(we, "bar failed")
	assert.True(t, errors.Cause(wme) == sql.ErrNoRows)

	fmt.Printf("data not found, %v\n", wme)
	fmt.Printf("%+v\n", wme)
}
