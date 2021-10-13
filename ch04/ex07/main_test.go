package main

import (
	"bytes"
	"log"
	"testing"
)

func TestReverse(t *testing.T) {
	actual := reverse(bytes.NewBufferString("あいう").Bytes())
	expected := bytes.NewBufferString("ういあ").Bytes()
	for i, _ := range actual {
		if actual[i] != expected[i] {
			log.Fatal("error")
		}
	}
}
