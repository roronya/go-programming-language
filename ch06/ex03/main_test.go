package main

import (
	"log"
	"testing"
)

func TestIntSet_IntersectWith(t *testing.T) {
	tests := []struct {
		s        []int
		t        []int
		expected []int
	}{
		{
			[]int{1, 2, 3},
			[]int{3, 4, 5},
			[]int{3},
		},
		{
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			[]int{},
		},
		{
			[]int{1, 2, 3, 5, 8},
			[]int{1, 2, 3, 5, 8},
			[]int{1, 2, 3, 5, 8},
		},
		{
			[]int{1, 2, 3, 5, 8},
			[]int{},
			[]int{},
		},
		{
			[]int{},
			[]int{1, 2, 3, 5, 8},
			[]int{},
		},
		{
			[]int{},
			[]int{},
			[]int{},
		},
	}

	for i, test := range tests {
		var s, t, expected IntSet
		s.AddAll(test.s...)
		t.AddAll(test.t...)
		expected.AddAll(test.expected...)

		s.IntersectWith(&t)

		if expected.String() != s.String() {
			log.Fatalf("case %d: expected %s, but got %s", i, expected.String(), s.String())
		}
	}
}

func TestIntSet_DifferenceWith(t *testing.T) {
	tests := []struct {
		s        []int
		t        []int
		expected []int
	}{
		{
			[]int{1, 2, 3},
			[]int{3, 4, 5},
			[]int{1, 2},
		},
		{
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			[]int{1, 2, 3},
		},
		{
			[]int{1, 2, 3, 5, 8},
			[]int{1, 2, 3, 5, 8},
			[]int{},
		},
		{
			[]int{1, 2, 3, 5, 8},
			[]int{},
			[]int{1, 2, 3, 5, 8},
		},
		{
			[]int{},
			[]int{1, 2, 3, 5, 8},
			[]int{},
		},
		{
			[]int{},
			[]int{},
			[]int{},
		},
	}

	for i, test := range tests {
		var s, t, expected IntSet
		s.AddAll(test.s...)
		t.AddAll(test.t...)
		expected.AddAll(test.expected...)

		s.DifferenceWith(&t)

		if expected.String() != s.String() {
			log.Fatalf("case %d: expected %s, but got %s", i, expected.String(), s.String())
		}
	}
}

func TestIntSet_SymmetricDifference(t *testing.T) {
	tests := []struct {
		s        []int
		t        []int
		expected []int
	}{
		{
			[]int{1, 2, 3},
			[]int{3, 4, 5},
			[]int{1, 2, 4, 5},
		},
		{
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			[]int{1, 2, 3, 5, 8},
			[]int{1, 2, 3, 5, 8},
			[]int{},
		},
		{
			[]int{1, 2, 3, 5, 8},
			[]int{},
			[]int{1, 2, 3, 5, 8},
		},
		{
			[]int{},
			[]int{1, 2, 3, 5, 8},
			[]int{1, 2, 3, 5, 8},
		},
		{
			[]int{},
			[]int{},
			[]int{},
		},
	}

	for i, test := range tests {
		var s, t, expected IntSet
		s.AddAll(test.s...)
		t.AddAll(test.t...)
		expected.AddAll(test.expected...)

		s.SymmetricDifference(&t)

		if expected.String() != s.String() {
			log.Fatalf("case %d: expected %s, but got %s", i, expected.String(), s.String())
		}
	}
}
