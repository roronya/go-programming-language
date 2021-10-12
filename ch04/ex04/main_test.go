package main

import (
	"log"
	"testing"
)

func TestRotate(t *testing.T) {
	actual := rotate([]int{1, 2, 3, 4, 5})
	expected := []int{2, 3, 4, 5, 1}
	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			log.Fatal("error")
		}
	}
}
