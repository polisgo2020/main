package sort

// go test -benchmem -bench .

import (
	"math/rand"
	"testing"
)

var (
	in []int
)

func init() {
	in = make([]int, 0, 1000)
	for i := 0; i < 1000; i++ {
		in = append(in, rand.Intn(9999))
	}
}

func BenchmarkBubble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bubblesort(in)
	}
}

func BenchmarkMerge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mergesort(in)
	}
}
