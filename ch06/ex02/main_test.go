package main

import (
	"log"
	"testing"
)

func TestIntSet_AddAll(t *testing.T) {
	tests := []struct {
		in       []int
		expected string
	}{
		{
			[]int{1, 2, 3, 5, 8},
			"{1 2 3 5 8}",
		},
		{
			[]int{},
			"{}",
		},
	}

	for _, test := range tests {
		var x IntSet
		x.AddAll(test.in...)
		if x.String() != test.expected {
			log.Fatalf("expected %s, but got %s\n", test.expected, x.String())
		}
	}

	// 何も渡さないケース
	var x IntSet
	x.AddAll()
	if x.String() != "{}" {
		log.Fatalf("expected {}, but got %s\n", x.String())
	}
}
