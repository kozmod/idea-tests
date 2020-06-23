package testify_test

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/mock"
	"testing"
)

type testInterface interface {
	Get(string, interface{}) error
}

type testMockedInterface struct {
	mock.Mock
}

func (m *testMockedInterface) Get(key string, data interface{}) error {
	args := m.Called(key, data)
	return args.Error(0)
}

type someStruct struct {
	Val string
}

const get = "Get"

func Test(t *testing.T) {
	mockInterface := new(testMockedInterface)
	mockInterface.On(get, "a", someStruct{"A"}).Return(errors.New("A error"))
	mockInterface.On(get, "a", someStruct{"B"}).Return(errors.New("B error"))

	if err := mockInterface.Get("a", someStruct{"A"}); err != nil {
		fmt.Println(err)
	}
	if err := mockInterface.Get("a", someStruct{"B"}); err != nil {
		fmt.Println(err)
	}
	mockInterface.AssertNumberOfCalls(t, get, 2)
}
