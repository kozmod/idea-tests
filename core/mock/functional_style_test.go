package mock

import (
	"github.com/magiconair/properties/assert"
	"strconv"
	"testing"
)

type testFuncType = func(int) string

var testFunc = func(i int) string {
	return strconv.Itoa(i)
}

func doSomeWork(i int, f testFuncType) string {
	i++
	return f(i)
}

func TestFunctionalStyle(t *testing.T) {
	res := doSomeWork(1, testFunc)
	assert.Equal(t, res, "2")
}

func TestFunctionalStyleMock(t *testing.T) {
	mockFunc := func(i int) string {
		assert.Equal(t, i, 2)
		return ""
	}
	doSomeWork(1, mockFunc)
}
