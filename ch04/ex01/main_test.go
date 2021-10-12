package main

import (
	"crypto/sha256"
	"log"
	"testing"
)

func TestDiffCount_同じ値(t *testing.T) {
	a := sha256.Sum256([]byte("x"))
	b := sha256.Sum256([]byte("x"))
	actual := diffCount(a, b)
	if actual != 0 {
		log.Fatal("error")
	}
}
