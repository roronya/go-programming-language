package popcount

import "testing"

func BenchmarkPopCount4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount4(0xffffffffffffffff)
	}
}

func BenchmarkPopCount4_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount4(0x5555555555555555)
	}
}
