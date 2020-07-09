package person

import (
	"reflect"
	"sync"
	"unsafe"
)

type emptyInterface struct {
	typ  *struct{}
	word unsafe.Pointer
}

var (
	offset1 uintptr
	offset2 uintptr
	offset3 uintptr
	p       Person
	t       = reflect.TypeOf(p)
	Pool    sync.Pool
)

func init() {
	offset1 = t.Field(1).Offset
	offset2 = t.Field(2).Offset
	offset3 = t.Field(3).Offset
	Pool.New = func() interface{} {
		return reflect.New(t)
	}
	for i := 0; i < 100; i++ {
		Pool.Put(reflect.New(t).Elem())
	}
}

type Person struct {
	Age   int
	Name  string
	Test1 string
	Test2 string
}

func New() interface{} {
	return &Person{
		Age:   30,
		Name:  "Kozmo",
		Test1: "test1",
		Test2: "test2",
	}
}

func NewUseReflect() interface{} {
	var p Person
	t := reflect.TypeOf(p)
	v := reflect.New(t)
	v.Elem().Field(0).Set(reflect.ValueOf(30))
	v.Elem().Field(1).Set(reflect.ValueOf("Kozmo"))
	v.Elem().Field(2).Set(reflect.ValueOf("test1"))
	v.Elem().Field(3).Set(reflect.ValueOf("test2"))
	return v.Interface()
}

//noinspection ALL
func NewQuickReflect() interface{} {
	v := reflect.New(t)

	p := v.Interface()
	ptr0 := uintptr((*emptyInterface)(unsafe.Pointer(&p)).word)
	ptr1 := ptr0 + offset1
	ptr2 := ptr0 + offset2
	ptr3 := ptr0 + offset3
	*((*int)(unsafe.Pointer(ptr0))) = 30
	*((*string)(unsafe.Pointer(ptr1))) = "Kozmo"
	*((*string)(unsafe.Pointer(ptr2))) = "test1"
	*((*string)(unsafe.Pointer(ptr3))) = "test2"
	return p
}

//noinspection ALL
func NewQuickReflectWithPool() interface{} {
	p := Pool.Get()
	ptr0 := uintptr((*emptyInterface)(unsafe.Pointer(&p)).word)
	ptr1 := ptr0 + offset1
	ptr2 := ptr0 + offset2
	ptr3 := ptr0 + offset3
	*((*int)(unsafe.Pointer(ptr0))) = 30
	*((*string)(unsafe.Pointer(ptr1))) = "Kozmo"
	*((*string)(unsafe.Pointer(ptr2))) = "test1"
	*((*string)(unsafe.Pointer(ptr3))) = "test2"
	return p
}
