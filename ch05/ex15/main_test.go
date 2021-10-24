package main

import (
	"log"
	"testing"
)

func TestMax(t *testing.T) {
	_, err := max()
	if err == nil {
		log.Fatalf("when no argument, err must not be nil\n")
	}

	tests := []struct {
		vals     []int
		expected int
	}{
		{[]int{1}, 1},
		{[]int{1, 2, 3, 4}, 4},
		{[]int{3, 4, 2, 1}, 4},
	}

	for _, test := range tests {
		actual, err := max(test.vals...)
		if err != nil {
			log.Fatalf("err must be nil, but got %s\n", err)
		}
		if actual != test.expected {
			log.Fatalf("expect %d, but got %d\n", test.expected, actual)
		}
	}
}

func TestMin(t *testing.T) {
	_, err := min()
	if err == nil {
		log.Fatalf("when no argument, err must not be nil\n")
	}

	tests := []struct {
		vals     []int
		expected int
	}{
		{[]int{1}, 1},
		{[]int{1, 2, 3, 4}, 1},
		{[]int{3, 4, 2, 1}, 1},
	}

	for _, test := range tests {
		actual, err := min(test.vals...)
		if err != nil {
			log.Fatalf("err must be nil, but got %s\n", err)
		}
		if actual != test.expected {
			log.Fatalf("expect %d, but got %d\n", test.expected, actual)
		}
	}
}

func TestMax2(t *testing.T) {
	actual := max2(1)
	if actual != 1 {
		log.Fatalf("expect 1 but got %d", actual)
	}

	tests := []struct {
		first    int
		vals     []int
		expected int
	}{
		{1, []int{}, 1},
		{1, []int{2, 3, 4}, 4},
		{3, []int{4, 2, 1}, 4},
	}

	for _, test := range tests {
		actual := max2(test.first, test.vals...)
		if actual != test.expected {
			log.Fatalf("expect %d, but got %d\n", test.expected, actual)
		}
	}
}
func TestMin2(t *testing.T) {
	actual := min2(1)
	if actual != 1 {
		log.Fatalf("expect 1 but got %d", actual)
	}

	tests := []struct {
		first    int
		vals     []int
		expected int
	}{
		{1, []int{}, 1},
		{1, []int{2, 3, 4}, 1},
		{3, []int{4, 2, 1}, 1},
	}

	for _, test := range tests {
		actual := min2(test.first, test.vals...)
		if actual != test.expected {
			log.Fatalf("expect %d, but got %d\n", test.expected, actual)
		}
	}
}
