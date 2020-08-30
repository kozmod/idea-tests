package internal

import (
	"github.com/kozmod/idea-tests/core/patterns/concurrency/pipeline/internal/pipe"
	"testing"
)

func BenchmarkWithoutPipelineModule(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addFoo(addQuoute(square(multiplyTwo(i))))
	}
}

func BenchmarkWithPipelineModule(b *testing.B) {
	outC := pipe.New(func(inC chan interface{}) {
		defer close(inC)
		for i := 0; i < b.N; i++ {
			inC <- i
		}
	}).
		Pipe(func(in interface{}) (interface{}, error) { return multiplyTwo(in.(int)), nil }).
		Pipe(func(in interface{}) (interface{}, error) { return square(in.(int)), nil }).
		Pipe(func(in interface{}) (interface{}, error) { return addQuoute(in.(int)), nil }).
		Pipe(func(in interface{}) (interface{}, error) { return addFoo(in.(string)), nil }).
		Merge()
	for result := range outC {
		_ = result
		// Do nothing, just for drain out from channels
	}
}
