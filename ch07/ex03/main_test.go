package main

import (
	"log"
	"testing"
)

func TestTree_String(t *testing.T) {
	tests := []struct {
		in       tree
		expected string
	}{
		{
			tree{
				0,
				&tree{1, nil, nil},
				&tree{2, nil, nil},
			},
			"[0, [1], [2]]",
		},
		{
			tree{
				0,
				&tree{1, &tree{2, nil, nil}, &tree{3, nil, nil}},
				&tree{4, &tree{5, nil, nil}, &tree{6, nil, nil}},
			},
			"[0, [1, [2], [3]], [4, [5], [6]]]",
		},
		{
			tree{},
			"[0]",
		},
	}

	for _, test := range tests {
		actual := test.in.String()
		if test.expected != actual {
			log.Fatalf("expected %s, but got %s\n", test.expected, actual)
		}
	}
}
