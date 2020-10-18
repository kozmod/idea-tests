package include

import "testing"

type someReturnDataA interface {
	float() float64
}

type someReturnDataB interface {
	float() float64
}

type someExecutorA interface {
	doWork(someReturnDataA)
}

type someExecutorB interface {
	doWork(someReturnDataB)
}

func returnFnA(se someExecutorA) someExecutorA {
	return se
}

func returnFnB(se someExecutorB) someExecutorB {
	return se
}

func TestDifInterfaces_1(t *testing.T) {
	var sea someExecutorA
	var seb someExecutorB
	var sa someReturnDataA
	var sb someReturnDataB

	sea.doWork(sa)
	sea.doWork(sb)
	seb.doWork(sa)
	seb.doWork(sb)

	returnFnA(sea)
	//returnFn(seb)// compilation error
}

func TestDifInterfaces_2(t *testing.T) {
	a := someReturnDataImpA{}
	ae := someExecutorImpA{}
	ae.doWork(a)
	returnFnA(ae)
	//returnFnB(ae)// compilation error

}

type someReturnDataImpA struct{}

func (s someReturnDataImpA) float() float64 {
	return 1.0
}

type someExecutorImpA struct{}

func (s someExecutorImpA) doWork(a someReturnDataA) {
}
