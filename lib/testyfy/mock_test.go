package testyfy_test

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

func Test(t *testing.T) {
	mockInterface := new(testMockedInterface)
	mockInterface.On("Get", "a", someStruct{"A"}).Return(errors.New("A error")).Once()
	mockInterface.On("Get", "a", someStruct{"B"}).Return(errors.New("B error")).Once()

	if err := mockInterface.Get("a", someStruct{"A"}); err != nil {
		fmt.Println(err)
	}
	if err := mockInterface.Get("a", someStruct{"B"}); err != nil {
		fmt.Println(err)
	}
}
