package popcount

func PopCount4(x uint64) int {
	y := x
	c := 0
	for y != 0 {
		y = y & (y - 1)
		c++
	}
	return c
}
