package popcount

import "testing"

var input uint64 = 0xffffffffffffffff

func BenchmarkPopCount2_6_2(b *testing.B, size int) {
	for i := 0; i < size; i++ {
		PopCount2_6_2(input)
	}
}

func Benchmark10PopCount2_6_2(b *testing.B) {
	BenchmarkPopCount2_6_2(b, 10)
}
