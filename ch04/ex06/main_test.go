package ex06

import (
	"bytes"
	"log"
	"testing"
)

func TestF(t *testing.T) {
	slice := bytes.NewBufferString("\t").Bytes()
	actual := f(slice)
	expected := bytes.NewBufferString(" ").Bytes()
	for i, _ := range actual {
		if actual[i] != expected[i] {
			log.Fatalf("error! %v, %v", actual, expected)
		}
	}
}
