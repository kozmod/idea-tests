package cpu_prof

import (
	"fmt"
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

func withCopy(s S) S {
	return s
}

func withPointer(s *S) *S {
	return s
}
func Benchmark_BeginsWithCopy(b *testing.B) {
	var s S
	for n := 0; n < b.N; n++ {
		s = byCopy()
	}
	_ = fmt.Sprintf("%v", s.a)
}

func Benchmark_BeginsWithArgCopy(b *testing.B) {
	var s S
	s = byCopy()
	for n := 0; n < b.N; n++ {
		//s = byCopy()
		s = withCopy(s)
	}
	_ = fmt.Sprintf("%v", s.a)
}

//noinspection ALL
func Benchmark_BeginsWithPointer(b *testing.B) {
	var s *S
	for n := 0; n < b.N; n++ {
		s = byPointer()
	}
	_ = fmt.Sprintf("%v", s.a)
}

func Benchmark_BeginsWithArgPointer(b *testing.B) {
	var s *S
	s = byPointer()
	for n := 0; n < b.N; n++ {
		s = withPointer(s)
	}
	_ = fmt.Sprintf("%v", s.a)
}
