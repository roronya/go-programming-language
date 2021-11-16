package main

import (
	"log"
	"sort"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		in       sort.Interface
		expected bool
	}{
		{sort.StringSlice{"a", "b", "c", "b", "a"}, true},
		{sort.StringSlice{"い", "ろ", "は", "ろ", "い"}, true},
		{sort.StringSlice{}, true},
		{sort.StringSlice{"a", "b"}, false},
		{sort.IntSlice{1, 2, 3, 2, 1}, true},
	}
	for i, t := range tests {
		actual := IsPalindrome(t.in)
		if t.expected != actual {
			log.Fatalf("case %d: expected %v, but got %v\n", i, t.expected, actual)
		}
	}
}
