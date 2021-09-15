package popcount

import (
	"testing"

	"github.com/roronya/go-programming-language/ch02/popcount"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0xffffffffffffffff)
	}
}

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount2(0xffffffffffffffff)
	}
}
