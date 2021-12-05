package main

import (
	"log"
	"testing"
)

func TestTopoSort_cyclic(t *testing.T) {
	asyclic := map[string]map[string]bool{
		"a": {"b": true},
		"b": {"c": true},
		"c": {"a": true},
	}
	actual, err := topoSort(asyclic)
	if err == nil {
		log.Fatalf("must be detected asyclic.\ngot = %#v\n", actual)
	}
}

func TestTopoSort_noncyclic(t *testing.T) {
	nonasyclic := map[string]map[string]bool{
		"a": {"b": true},
		"b": {"c": true},
		"c": {"d": true},
	}
	actual, err := topoSort(nonasyclic)
	if err != nil {
		log.Fatalf("must not be detected asyclic.\ngot = %#v\n", actual)
	}
}

func TestTopoSort_prereqs(t *testing.T) {
	actual, err := topoSort(prereqs)
	if err == nil {
		log.Fatalf("must be detected asyclic.\ngot = %#v\n", actual)
	}
}
