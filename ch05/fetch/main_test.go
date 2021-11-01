package main

import (
	"fmt"
	_io "io"
	"testing"
)

type OS struct {
}

func (os *OS) Create(filename string) error {
	if filename == "error" {
		return fmt.Errorf("error in os.Create()")
	}
	return nil
}

type IO struct{}

func (io *IO) Copy(src _io.Reader, dst _io.Writer) error {
	return fmt.Errorf("error in io.Copy()")
}

var os OS
var io IO

func TestFetch(t *testing.T) {
	fetch("https://example.com")
}
