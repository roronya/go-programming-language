package main

import (
	"fmt"
	"log"
	"testing"
)

func TestWordCounter_Write(t *testing.T) {
	tests := []struct {
		in       string
		expected WordCounter
	}{
		{
			"hello world",
			2,
		},
		{
			"",
			0,
		},
		{
			"こんにちは　世界",
			2,
		},
	}

	for _, test := range tests {
		var actual WordCounter
		fmt.Fprintf(&actual, test.in)
		if test.expected != actual {
			log.Fatalf("expected %d, but got %d", test.expected, actual)
		}
	}

}

func TestNewlineCounter_Write(t *testing.T) {
	tests := []struct {
		in       string
		expected NewlineCounter
	}{
		{
			"hello world",
			1,
		},
		{
			"hello\nworld",
			2,
		},
		{
			"",
			0,
		},
		{
			"こんにちは\n世界",
			2,
		},
	}

	for _, test := range tests {
		var actual NewlineCounter
		fmt.Fprintf(&actual, test.in)
		if test.expected != actual {
			log.Fatalf("expected %d, but got %d", test.expected, actual)
		}
	}

}
