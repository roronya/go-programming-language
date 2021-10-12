package main

import (
	"crypto/sha256"
	"fmt"
	"math/bits"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	d := diffCount(c1, c2)
	fmt.Printf("%d", d)
}

func diffCount(a [32]uint8, b [32]uint8) int {
	d := 0
	for i := 0; i < len(a); i++ {
		d += bits.OnesCount8(a[i] ^ b[i])
	}
	return d
}
