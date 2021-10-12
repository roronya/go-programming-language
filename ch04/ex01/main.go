package main

import (
	"crypto/sha256"
	"fmt"
	"math/bits"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	d := 0
	for i := 0; i < len(c1); i++ {
		d += bits.OnesCount8(c1[i] ^ c2[i])
	}
	fmt.Printf("%d", d)
}
