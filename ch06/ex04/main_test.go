package main

import (
	"log"
	"testing"
)

func TestIntSet_IntersectWith(t *testing.T) {
	tests := []struct {
		s []int
	}{
		{
			[]int{1, 2, 3},
		},
		{
			[]int{},
		},
	}

	for i, test := range tests {
		var s IntSet
		s.AddAll(test.s...)

		actual := s.Elems()

		for j, _ := range actual {
			if test.s[j] != actual[j] {
				log.Fatalf("case %d: error\n", i)
			}
		}
	}
}
