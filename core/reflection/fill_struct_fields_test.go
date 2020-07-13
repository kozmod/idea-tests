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

type injectedInterfaceImplB struct {
	Val string
}

func (iii injectedInterfaceImplB) getVal() string {
	return iii.Val
}

type testReflectionStruct struct {
	Val string
	I   injectedInterface
}

func TestFillStructFieldFromInstance(t *testing.T) {
	injection := injectedInterfaceImplB{b}
	obj := testReflectionStruct{}
	val := reflect.ValueOf(&obj)
	fmt.Println(val)
	for i := 0; i < val.Elem().NumField(); i++ {
		switch val.Elem().Field(i).Type() {
		case typeOfInterface:
			val.Elem().Field(i).Set(reflect.ValueOf(&injection))
		case typeOfString:
			val.Elem().Field(i).SetString(a)

		}
	}
	fmt.Println(val)
	assert.Equal(t, obj.I.getVal(), b)
	assert.Equal(t, obj.Val, a, obj.Val)
}

func TestFillStructFieldFromInstancesMap(t *testing.T) {
	m := make(map[string]interface{})
	m[b] = injectedInterfaceImplB{b}

	obj := testReflectionStruct{}
	val := reflect.ValueOf(&obj)
	fmt.Println(val)
	for i := 0; i < val.Elem().NumField(); i++ {
		switch val.Elem().Field(i).Type() {
		case typeOfInterface:
			injection := m[b]
			val.Elem().Field(i).Set(reflect.ValueOf(injection)) //dif from TestFillStructFieldFromInstance
		case typeOfString:
			val.Elem().Field(i).SetString(a)

		}
	}
	fmt.Println(val)
	assert.Equal(t, obj.I.getVal(), b)
	assert.Equal(t, obj.Val, a, obj.Val)
}

func TestFillStructFieldFromInstancesMap_2(t *testing.T) {
	m := make(map[string]interface{})

	iib := injectedInterfaceImplB{}

	assert.Equal(t, iib.getVal(), "")

	vb := reflect.ValueOf(&iib)
	vb.Elem().Field(0).Set(reflect.ValueOf(b))
	m[b] = iib

	assert.Equal(t, iib.getVal(), b)

	obj := testReflectionStruct{}
	val := reflect.ValueOf(&obj)
	fmt.Println(val)
	for i := 0; i < val.Elem().NumField(); i++ {
		switch val.Elem().Field(i).Type() {
		case typeOfInterface:
			injection := m[b]
			val.Elem().Field(i).Set(reflect.ValueOf(injection))
		case typeOfString:
			val.Elem().Field(i).SetString(a)

		}
	}
	fmt.Println(val)
	assert.Equal(t, obj.I.getVal(), b)
	assert.Equal(t, obj.Val, a, obj.Val)
}
