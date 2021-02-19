package search

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

var testCases []testCase

type testCase struct {
	in  []int
	exp int
}

func init() {
	const quantity = 10
	cases := make([]testCase, 0, quantity)
	set := make(map[int]struct{})
	for i := 0; i < quantity; i++ {
		input := make([]int, 0, quantity)
		var max *int
		for j := 0; j < quantity; j++ {
			for {
				randi := rand.Intn(100)
				_, ok := set[randi]
				if ok {
					continue
				}
				set[randi] = struct{}{}
				if max == nil || *max < randi {
					max = &randi
				}
				input = append(input, randi)
				break
			}
		}
		if max == nil || len(input) < 1 {
			i := 0
			max = &i
			input = []int{0}
		}
		cases = append(cases, testCase{in: input, exp: *max})
	}
	testCases = cases
}

func BenchmarkReqMax(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, cs := range testCases {
			reqMax(cs.in)
		}
	}
}

func BenchmarkLinearMax(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, cs := range testCases {
			linearMax(cs.in)
		}
	}
}

func TestReqMax(t *testing.T) {
	for i, cs := range testCases {
		res := reqMax(cs.in)
		assert.Equal(t, cs.exp, res, fmt.Sprintf("exp: %d, actual:%d, case:%d", cs.exp, res, i))
	}
}

func TestLinearMax(t *testing.T) {
	for i, cs := range testCases {
		res := linearMax(cs.in)
		assert.Equal(t, cs.exp, res, fmt.Sprintf("exp: %d, actual:%d, case:%d", cs.exp, res, i))
	}
}

func reqMax(in []int) int {
	if len(in) < 2 {
		return in[0]
	}
	l := reqMax(in[:len(in)/2])
	r := reqMax(in[len(in)/2:])
	if l > r {
		return l
	}
	return r
}

func linearMax(in []int) int {
	var max *int
	for _, v := range in {
		if max == nil || *max < v {
			max = &v
		}
	}
	if max == nil {
		return 0
	}
	return *max
}
