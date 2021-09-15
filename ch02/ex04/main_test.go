package popcount

import (
	"testing"
)

func BenchmarkPopCount3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount3(0xffffffffffffffff)
	}
}
