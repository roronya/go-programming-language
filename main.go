package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := make(map[string]string)
	x := a["a"]
	fmt.Println(strconv.Atoi(x))
}
