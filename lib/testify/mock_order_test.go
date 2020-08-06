package testify

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
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

	c := &Composition{amock, bmock}
	err := c.Apply()
	require.NoError(t, err)
}

const (
	executeA = "ExecuteA"
	executeB = "ExecuteB"
)

type SpyA struct {
	Calls *[]string
}

func (s *SpyA) Execute() {
	*s.Calls = append(*s.Calls, executeA)
}

type SpyB struct {
	Calls *[]string
}

func (s *SpyB) Execute() {
	*s.Calls = append(*s.Calls, executeB)
}

func TestOrderOfMocks_Spy(t *testing.T) {
	calls := &[]string{}
	sa := &SpyA{calls}
	sb := &SpyB{calls}
	c := &Composition{
		a: sa,
		b: sb,
	}
	err := c.Apply()
	require.NoError(t, err)
	expected := []string{
		0: executeA,
		1: executeB,
	}
	for i, val := range *calls {
		if expected[i] != val {
			t.Fail()
		}
	}
}
