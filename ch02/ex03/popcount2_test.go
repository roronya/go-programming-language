package popcount2

import (
	"testing"

	"github.com/roronya/go-programming-language/ch02/popcount"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(1000000)
	}
}

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(1000000)
	}
}
