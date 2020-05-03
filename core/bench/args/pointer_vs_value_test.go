package args

import (
	"fmt"
	"os"
	"runtime/trace"
	"testing"
)

type S struct {
	a, b, c int64
	d, e, f string
	g, h, i float64
}

func byCopy() S {
	return S{
		a: 1, b: 1, c: 1,
		e: "foo", f: "foo",
		g: 1.0, h: 1.0, i: 1.0,
	}
}

func byPointer() *S {
	return &S{
		a: 1, b: 1, c: 1,
		e: "foo", f: "foo",
		g: 1.0, h: 1.0, i: 1.0,
	}
}

/*
go test ./... -bench=BenchmarkMemoryStack -benchmem -run=^$ -count=10 2>&1 | tee out/stack.txt && benchstat out/stack.txt 2>&1 | tee out/stack-bstat.txt
*/
func BenchmarkMemoryStack(b *testing.B) {
	var s S

	f, err := os.Create("out/stack.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		s = byCopy()
		//byCopyArg(s)
	}

	trace.Stop()

	b.StopTimer()

	_ = fmt.Sprintf("%v", s.a)
}

/*
go test ./... -bench=BenchmarkMemoryHeap -benchmem -run=^$ -count=10 2>&1 | tee out/head.txt && benchstat out/head.txt 2>&1 | tee out/head-bstat.txt
*/
func BenchmarkMemoryHeap(b *testing.B) {
	var s *S

	f, err := os.Create("out/heap.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		s = byPointer()
		//byPointerArg(s)
	}

	trace.Stop()

	b.StopTimer()

	_ = fmt.Sprintf("%v", s.a)
}
