package regex

import (
	"testing"
)

const cap = 1_000

func BenchmarkEmptyAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 0)
		for i := 0; i < cap; i++ {
			s = append(s, i)
		}
	}
}

func BenchmarkPreallocateAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 0, cap)
		for i := 0; i < cap; i++ {
			s = append(s, i)
		}
	}
}
