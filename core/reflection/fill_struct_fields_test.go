package reflection_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/magiconair/properties/assert"
)

const (
	a = "valA"
	b = "valB"
)

var (
	typeOfInterface = reflect.TypeOf((*injectedInterface)(nil)).Elem()
	typeOfString    = reflect.TypeOf("")
)

type injectedInterface interface {
	getVal() string
}

type injectedInterfaceImplA struct{}

func (iii injectedInterfaceImplA) getVal() string {
	return a
}

type injectedInterfaceImplB struct{}

func (iii injectedInterfaceImplB) getVal() string {
	return b
}

type testReflectionStruct struct {
	Val string
	I   injectedInterface
}

func Test(t *testing.T) {
	injection := injectedInterfaceImplB{}
	obj := testReflectionStruct{}
	val := reflect.ValueOf(&obj)
	fmt.Println(val)
	for i := 0; i < val.Elem().NumField(); i++ {

		switch val.Elem().Field(i).Type() {
		case typeOfInterface:
			rv := reflect.ValueOf(injection)
			val.Elem().Field(i).Set(rv)
		case typeOfString:
			val.Elem().Field(i).SetString(a)

		}
	}
	assert.Equal(t, b, obj.I.getVal())
	assert.Equal(t, a, obj.Val)
}
