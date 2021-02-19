package search

import (
	"fmt"
	"github.com/kozmod/idea-tests/algorithms/classic/search"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

var testCasesSift []testCase

func init() {
	const quantity = 10
	cases := make([]testCase, 0, quantity)
	for i := 0; i < quantity; i++ {
		input := make([]int, quantity)
		sift := rand.Intn(quantity - 1)
		val := 0
		for j := sift; j < len(input); j++ {
			input[j] = val
			val++
		}
		for j := 0; j < sift; j++ {
			input[j] = val
			val++
		}
		cases = append(cases, testCase{in: input, exp: quantity - 1})
	}
	testCasesSift = cases
}

func BenchmarkShiftedReqMax(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, cs := range testCasesSift {
			reqMax(cs.in)
		}
	}
}

func BenchmarkShiftedLinearMax(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, cs := range testCasesSift {
			linearMax(cs.in)
		}
	}
}

func BenchmarkShiftedReqMax_2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, cs := range testCasesSift {
			_, _ = search.SearchSiftedMax(cs.in)
		}
	}
}

func TestSearchSiftedMax(t *testing.T) {
	for i, cs := range testCasesSift {
		res, _ := search.SearchSiftedMax(cs.in)
		assert.Equal(t, cs.exp, res, fmt.Sprintf("exp: %d, actual:%d, case:%d", cs.exp, res, i))
	}
}
