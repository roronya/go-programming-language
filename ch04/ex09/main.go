package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)

	counts := make(map[string]int)
	for input.Scan() {
		counts[input.Text()]++
	}

	for k, v := range counts {
		fmt.Printf("%s:%d\n", k, v)
	}
}
