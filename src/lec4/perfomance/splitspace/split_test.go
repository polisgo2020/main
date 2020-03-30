package splitspace

// go test -benchmem -bench .

import (
	"testing"
)

func BenchmarkTrim(b *testing.B) {
	for i := 0; i < b.N; i++ {
		splitTrim(in)
	}
}

func BenchmarkRegexp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tokenize(in)
	}
}

func BenchmarkRegexpGlobal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tokenizeGlobal(in)
	}
}
