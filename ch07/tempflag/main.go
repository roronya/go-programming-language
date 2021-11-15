package main

import (
	"flag"
	"fmt"

	"github.com/roronya/go-programming-language/ch07/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
