package main

import (
	"bytes"
	"testing"
)

func TestCoutingWriter(t *testing.T) {
	in := bytes.NewBuffer([]byte{})
	actualWriter, actualPtr := CoutingWriter(in)
}
