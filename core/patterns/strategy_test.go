package patterns

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

//OOP realization
type Operator interface {
	Apply(int, int) int
}

type Addition struct{}

func (Addition) Apply(lval, rval int) int {
	return lval + rval
}

type Multiplication struct{}

func (Multiplication) Apply(lval, rval int) int {
	return lval * rval
}

type Operation struct {
	Operator Operator
}

func (o *Operation) Operate(leftValue, rightValue int) int {
	return o.Operator.Apply(leftValue, rightValue)
}

func TestStrategyPattern(t *testing.T) {
	mult := Operation{Multiplication{}}
	assert.Equal(t, 15, mult.Operate(3, 5))

	add := Operation{Addition{}}
	assert.Equal(t, 8, add.Operate(3, 5))
}

//Functional realization
type FOperation struct {
	operation func(int, int) int
}

func (o *FOperation) Operate(leftValue, rightValue int) int {
	return o.operation(leftValue, rightValue)
}

func multiply(lval, rval int) int {
	return lval * rval
}

func add(lval, rval int) int {
	return lval + rval
}

func TestStrategyPatternFunctionalStyle(t *testing.T) {
	mult := FOperation{multiply}
	assert.Equal(t, 15, mult.Operate(3, 5))

	add := FOperation{add}
	assert.Equal(t, 8, add.Operate(3, 5))
}
