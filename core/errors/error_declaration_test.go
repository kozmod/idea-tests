package errors

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

var varError = errors.New("some error")

func SomeError() error {
	return varError
}

//goland:noinspection ALL
func TestErrorAsConstant(t *testing.T) {
	sql.ErrConnDone = errors.New("{wtf}")
	fmt.Println(sql.ErrConnDone)
	assert.True(t, errors.Is(SomeError(), varError))
	assert.True(t, errors.Is(fmt.Errorf("E2: %w", SomeError()), varError))
	assert.True(t, errors.Is(errors.Wrap(SomeError(), "wrap"), varError))
	assert.True(t, errors.Is(errors.WithMessage(SomeError(), "with msg"), varError))
	assert.True(t, errors.Is(errors.WithStack(SomeError()), varError))
	assert.True(t, errors.Is(errors.WithStack(SomeError()), varError))

	sql.ErrConnDone = varError
	assert.True(t, errors.Is(SomeError(), sql.ErrConnDone))
	assert.True(t, errors.Is(fmt.Errorf("E2: %w", SomeError()), sql.ErrConnDone))
	assert.True(t, errors.Is(errors.Wrap(SomeError(), "wrap"), sql.ErrConnDone))
	assert.True(t, errors.Is(errors.WithMessage(SomeError(), "with msg"), sql.ErrConnDone))
	assert.True(t, errors.Is(errors.WithStack(SomeError()), sql.ErrConnDone))
	assert.True(t, errors.Is(errors.WithStack(SomeError()), sql.ErrConnDone))

}
