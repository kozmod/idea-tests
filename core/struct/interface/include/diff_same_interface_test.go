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

func returnFn(se someExecutorA) someExecutorA {
	return se
}

func Test(t *testing.T) {
	var sea someExecutorA
	var seb someExecutorB
	var sa someReturnDataA
	var sb someReturnDataB

	sea.doWork(sa)
	sea.doWork(sb)
	seb.doWork(sa)
	seb.doWork(sb)

	returnFn(sea)
	//returnFn(seb)// compilation error
}
