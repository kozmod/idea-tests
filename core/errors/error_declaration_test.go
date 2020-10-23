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
	assert.False(t, errors.Is(nil, sql.ErrConnDone))
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

func TestWrapSqlError(t *testing.T) {
	sql.ErrConnDone = errors.New("{wtf}")
	fmt.Println(errors.Wrap(sql.ErrConnDone, "wrapped error"))
	fmt.Println(errors.WithMessage(sql.ErrConnDone, "wrapped error"))
	fmt.Println(fmt.Errorf("wrapped error: %w", sql.ErrConnDone))
}

func Test(t *testing.T) {
	result, err := caller1()
	if err != nil {
		fmt.Printf("%+v\n", errors.WithStack(err))
		return
	}
	fmt.Println("Result: ", result)
}

func caller1() (int, error) {
	err := caller2()
	if err != nil {
		return 0, err
	}
	return 1, nil
}

func caller2() error {
	err := caller3()
	if err != nil {
		return err
	}
	return nil
}

func caller3() error {
	return errors.New("failed")
}
