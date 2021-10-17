package main

import (
	"bufio"
	"bytes"

	"golang.org/x/net/html"
)

func countWordsAndImages(n *html.Node) (words, images int) {
	q := []*html.Node{n}
	var top *html.Node
	for !isEmpty(q) {
		q, top = pop(q)
		for c := top.FirstChild; c != nil; c = c.NextSibling {
			q = push(q, c)
		}

		if top.Type == html.TextNode {
			bufferString := bytes.NewBufferString(top.Data)
			input := bufio.NewScanner(bufferString)
			input.Split(bufio.ScanWords)
			for input.Scan() {
				words++
			}
		}

		if top.Type == html.ElementNode && top.Data == "img" {
			images++
		}
	}
	return words, images
}

func push(q []*html.Node, e *html.Node) []*html.Node {
	return append(q, e)
}

func pop(q []*html.Node) ([]*html.Node, *html.Node) {
	if len(q) == 0 {
		return q, nil
	}
	top := q[0]
	for i := 0; i < len(q)-1; i++ {
		q[i] = q[i+1]
	}
	return q[:len(q)-1], top
}

func isEmpty(q []*html.Node) bool {
	return len(q) == 0
}
