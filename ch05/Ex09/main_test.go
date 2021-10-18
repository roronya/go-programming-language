package main

import (
	"log"
	"testing"
)

func TestExpand(t *testing.T) {
	fooToBar := func(s string) string {
		if s == "foo" {
			return "bar"
		}
		return s
	}
	jToE := func(s string) string {
		if s == "こんにちは" {
			return "Hello"
		}
		return s
	}

	tests := []struct {
		in       string
		f        func(s string) string
		expected string
	}{
		{"$foo", fooToBar, "bar"},
		{"this is $foo and $foo", fooToBar, "this is bar and bar"},
		{"this is $", fooToBar, "this is "},
		{"$こんにちは　世界", jToE, "Hello 世界"},
	}
	for _, test := range tests {
		actual := expand(test.in, test.f)
		if actual != test.expected {
			log.Fatalf("%s is expected %s", actual, test.expected)
		}
	}
}
