package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	numberCounts := make(map[rune]int) // 数字の数
	letterCounts := make(map[rune]int) // 文字の数
	var utflen [utf8.UTFMax + 1]int    // UTF-8エンコーディングの長さの数
	invalid := 0                       // 不正なUTF-8文字の数

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // rune, nbytes, errorを返す
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			letterCounts[r]++
		}
		if unicode.IsNumber(r) {
			numberCounts[r]++

		}
		utflen[n]++
	}
	fmt.Printf("rune\tnumber count\n")
	for c, n := range numberCounts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("rune\tletter count\n")
	for c, n := range letterCounts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}

}
