package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // intをByteCounterへ変換
	return len(p), nil
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	buf := bytes.NewBuffer(p)
	input := bufio.NewScanner(buf)
	input.Split(bufio.ScanWords)

	wc := 0
	for input.Scan() {
		wc++
	}
	*c += WordCounter(wc)
	return wc, nil
}

type NewlineCounter int

func (c *NewlineCounter) Write(p []byte) (int, error) {
	buf := bytes.NewBuffer(p)
	input := bufio.NewScanner(buf)
	input.Split(bufio.ScanLines)

	nc := 0
	for input.Scan() {
		nc++
	}
	*c += NewlineCounter(nc)
	return nc, nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0
	name := "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}
