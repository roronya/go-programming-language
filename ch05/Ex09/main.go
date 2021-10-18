package main

import (
	"bufio"
	"bytes"
	"strings"
)

func expand(s string, f func(string) string) string {
	buf := bytes.NewBufferString(s)
	input := bufio.NewScanner(buf)
	input.Split(bufio.ScanWords)
	var result []string
	for input.Scan() {
		w := input.Text()
		if strings.HasPrefix(w, "$") {
			result = append(result, f(strings.TrimLeft(w, "$")))
		} else {
			result = append(result, w)
		}
	}
	return strings.Join(result, " ")
}
