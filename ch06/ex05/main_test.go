package main

import (
	"log"
	"testing"
)

func TestIntSet_BITWIDTH(t *testing.T) {
	if BITWIDTH != 64 {
		log.Fatalf("expected 64, but got %d\n", BITWIDTH)
	}
}
