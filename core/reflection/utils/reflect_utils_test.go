package utils_test

import (
	"reflect"
	"testing"

	"github.com/kozmod/idea-tests/core/reflection/utils"
	"github.com/magiconair/properties/assert"
)

const testVal = "testVal"

type Foo struct {
	unexportedField string
}

func TestGetUnexportedFields(t *testing.T) {
	foo := &Foo{testVal}
	field := utils.GetUnexportedField(reflect.ValueOf(foo).Elem().FieldByName("unexportedField"))
	assert.Equal(t, testVal, field)
}

func TestSetUnexportedField(t *testing.T) {
	foo := &Foo{"some"}
	utils.SetUnexportedField(reflect.ValueOf(foo).Elem().FieldByName("unexportedField"), testVal)
	assert.Equal(t, testVal, foo.unexportedField)
}
