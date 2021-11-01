package main

import (
	"fmt"
	"log"
)

func main() {
	err := f()
	log.Fatal(err)
}

func f() error {
	defer g()
	return fmt.Errorf("error in f()")
}

func g() error {
	return fmt.Errorf("error in defer")
}
