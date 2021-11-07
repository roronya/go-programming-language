package main

import (
	"log"
	"os"
	"testing"
)

func TestFind(t *testing.T) {
	file0, _ := os.Create("test.txt")
	defer file0.Close()
	defer os.Remove("test.txt")

	os.MkdirAll("a/b/c/", 0777)
	file1, _ := os.Create("a/b/c/test.txt")
	defer file1.Close()
	defer os.RemoveAll("a")

	actuals := find("./", "test.txt")
	if len(actuals) != 2 {
		log.Fatalf("len(actual) is expected 2, but got %d\n", len(actuals))
	}
	for _, actual := range actuals {
		if !("test.txt" == actual || "a/b/c/test.txt" == actual) {
			log.Fatalf("unexpected result, got %s\n", actual)
		}
	}
}
