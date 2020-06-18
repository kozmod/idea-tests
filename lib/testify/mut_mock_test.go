package testify_test

import (
	"fmt"
	"github.com/stretchr/testify/mock"
	"testing"
)

type testMut interface {
	doSth()
}

type testMutImpl struct {
	val string
}

func (tm *testMutImpl) doSth() {
}

type testWorker interface {
	Get(testMut) error
}

type testMockedWorker struct {
	mock.Mock
}

func (m *testMockedWorker) Get(tm testMut) error {
	args := m.Called(tm)
	return args.Error(0)
}

//todo finish test
func TestMut(t *testing.T) {
	mockInterface := new(testMockedWorker)
	//mockInterface.On("Get", &testMutImpl{"tm"}).Return(errors.New("A error")).Once()
	mockInterface.On("Get", &testMutImpl{"tm"}).Run(func(args mock.Arguments) {
		i := args.Get(0)
		fmt.Printf("%p,%v\n", i, i)
		i = &testMutImpl{"XXXX"}
		//reflect.ValueOf(i).SetPointer(unsafe.Pointer(&testMutImpl{"XXXX"}))
		fmt.Printf("%p,%v\n", i, i)
	}).Return(nil)

	x := &testMutImpl{"tm"}
	fmt.Printf("%p,%v\n", x, x)
	if err := mockInterface.Get(x); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%p,%v\n", x, x)
}
