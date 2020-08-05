package testify

import (
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

type InterdaceA interface {
	Execute()
}

type InterdaceB interface {
	Execute()
}
type Composition struct {
	a InterdaceA
	b InterdaceB
}

func (c *Composition) Apply() error {
	c.a.Execute() // how to check that "a" execute before "b"
	c.b.Execute()
	return nil
}

func New(a InterdaceA, b InterdaceB) *Composition {
	return &Composition{a, b}
}

const Execute = "Execute"

type AMock struct {
	mock.Mock
}

func (m *AMock) Execute() {
	m.Called()
}

type BMock struct {
	mock.Mock
}

func (m *BMock) Execute() {
	m.Called()
}

func TestOrderOfMocks(t *testing.T) {
	order := 0
	amock := new(AMock)
	amock.On(Execute).Run(func(args mock.Arguments) {
		if order++; order != 1 {
			t.Fail()
		}
	})

	bmock := new(BMock)
	bmock.On(Execute).Run(func(args mock.Arguments) {
		if order++; order != 2 {
			t.Fail()
		}
	})

	c := New(amock, bmock)
	err := c.Apply()
	require.NoError(t, err)
}
