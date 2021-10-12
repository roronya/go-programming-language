package main

import (
	"log"
	"testing"
)

func TestReverse(t *testing.T) {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(&a)

	if a != [...]int{5, 4, 3, 2, 1, 0} {
		log.Fatal("error!")
	}
}
