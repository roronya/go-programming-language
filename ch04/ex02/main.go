package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	var f1, f2 bool
	flag.BoolVar(&f1, "sha384", false, "sha384のハッシュ値を表示します")
	flag.BoolVar(&f2, "sha512", false, "sha512のハッシュ値を表示します")
	flag.Parse()
	d := []byte(flag.Arg(0))
	if f1 {
		fmt.Printf("%x", sha512.Sum384(d))
	} else if f2 {
		fmt.Printf("%x", sha512.Sum512(d))
	} else {
		fmt.Printf("%x", sha256.Sum256(d))
	}
}
