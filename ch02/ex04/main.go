package popcount

func PopCount3(x uint64) int {
	y := 0
	for i := 0; i < 64; i++ {
		y += int(x >> i & 1)
	}
	return y
}
