package regex

import (
	"regexp"
	"strings"
	"testing"
)

var (
	value = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36"
	rgexp = regexp.MustCompile("AppleWebKit")
	kit   = "AppleWebKit"
)

func BenchmarkRegExp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = regexp.MatchString(kit, value)
	}
}

func BenchmarkRegCompiled(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = rgexp.MatchString(value)
	}
}

func BenchmarkStringContains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strings.Contains(value, kit)
	}
}
