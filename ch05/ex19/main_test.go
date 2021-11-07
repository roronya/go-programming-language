package main

import (
	"log"
	"testing"
)

func TestF(t *testing.T) {
	actual := f()
	if actual != "recovered" {
		log.Fatalf("expected \"recovered\", but got %s\n", actual)
	}
}
