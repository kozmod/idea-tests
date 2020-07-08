package improve_speed_test

import (
	"testing"

	"github.com/kozmod/idea-tests/core/reflection/bech/improve_speed/include/person"
)

func BenchmarkNew(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		person.New()
	}
}

func BenchmarkNewUseReflect(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		person.NewUseReflect()
	}
}

func BenchmarkNewQuickReflect(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		person.NewQuickReflect()
	}
}

func BenchmarkNewQuickReflectLocalVar(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		person.NewQuickReflect()
	}
}

func BenchmarkQuickReflectWithPool(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		obj := person.NewQuickReflectWithPool()
		person.Pool.Put(obj)
	}
}
