package main

import (
	"log"
	"testing"
)

func TestIntSet_Len(t *testing.T) {
	tests := []struct {
		in       []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{1, 3, 5}, 3},
		{[]int{1, 1, 1}, 1},
		{[]int{100, 200, 300}, 3},
	}

	for _, test := range tests {
		var x IntSet
		for _, i := range test.in {
			x.Add(i)
		}
		actual := x.Len()
		if test.expected != actual {
			log.Fatalf("expected %d, but got %d\n", test.expected, actual)
		}
	}
}

func TestIntSet_Remove(t *testing.T) {
	tests := []struct {
		init     []int
		in       int
		expected []int
	}{
		{[]int{1, 2, 3, 5, 8}, 1, []int{2, 3, 5, 8}},
		{[]int{}, 1, []int{}},
		{[]int{1, 2, 3, 5, 8}, 100, []int{1, 2, 3, 5, 8}},
	}

	for _, test := range tests {
		var actual IntSet
		for _, i := range test.init {
			actual.Add(i)
		}
		actual.Remove(test.in)

		var expected IntSet
		for _, i := range test.expected {
			expected.Add(i)
		}
		if expected.String() != actual.String() {
			log.Fatalf("expected %s, but got %s\n", expected.String(), actual.String())
		}
	}
}

func TestIntSet_Clear(t *testing.T) {
	tests := []struct {
		init []int
	}{
		{[]int{1, 2, 3, 5, 8}},
		{[]int{}},
	}

	for _, test := range tests {
		var actual IntSet
		for _, i := range test.init {
			actual.Add(i)
		}
		actual.Clear()

		if actual.Len() != 0 {
			log.Fatalf("expected cleared, but got %s\n", actual.String())
		}
		actual.Add(0) // clear後に値を追加できるか
		if actual.Len() != 1 {
			log.Fatalf("cannot add value after clearing")
		}
	}
}

func TestIntSet_Copy(t *testing.T) {
	tests := []struct {
		init []int
	}{
		{[]int{1, 2, 3, 5, 8}},
		{[]int{}},
	}

	for _, test := range tests {
		var actual IntSet
		for _, i := range test.init {
			actual.Add(i)
		}
		expected := actual.Copy()

		// アドレスが異なること
		if expected == &actual {
			log.Fatalf("expected == &actual")
		}
		if &expected.words == &actual.words {
			log.Fatalf("&expected.words == &actual.words")
		}
		// 入っている値は同じであること
		if expected.String() != actual.String() {
			log.Fatalf("expected %s, but got %s\n", expected.String(), actual.String())
		}
	}
}
