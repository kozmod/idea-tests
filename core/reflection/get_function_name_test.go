package reflection_test

import (
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func foo() {
}

func TestGetFunctionNameTest(t *testing.T) {
	fullName := runtime.FuncForPC(reflect.ValueOf(foo).Pointer()).Name()
	fullNameSlice := strings.Split(fullName, ".")
	name := fullNameSlice[len(fullNameSlice)-1]
	assert.Equal(t, "foo", name)
}
