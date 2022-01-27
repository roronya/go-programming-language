package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount2_6_2(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCount2_4(x uint64) int {
	y := 0
	for i := 0; i < 64; i++ {
		y += int(x >> i & 1)
	}
	return y
}

func PopCount2_5(x uint64) int {
	y := x
	c := 0
	for y != 0 {
		y = y & (y - 1)
		c++
	}
	return c
}
