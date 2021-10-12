package main

import (
	"log"
	"testing"
)

func TestF(t *testing.T) {
	actual := f([]string{"a", "a", "b", "c", "c", "a"})
	expected := []string{"a", "b", "c", "a"}
	if len(actual) != len(expected) {
		log.Fatal("error")
	}
	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			log.Fatal("error")
		}
	}
}
