package main

import (
	"fmt"
	"testing"
)

func TestFetch(t *testing.T) {
	filename, n, err := fetch("https://example.com")
	fmt.Printf("%s %d %s", filename, n, err)
}
