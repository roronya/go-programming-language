package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stdin, "findlinnks1: %v\n", err)
		os.Exit(1)
	}
	visit(doc)
}

func visit(n *html.Node) {
	q := []*html.Node{n}
	var top *html.Node // 宣言しておかないと26Lでshadowingされてしまう
	for !isEmpty(q) {
		q, top = pop(q)
		for c := top.FirstChild; c != nil; c = c.NextSibling {
			q = push(q, c)
		}

		if top.Type == html.TextNode && top.Parent.Data != "script" && top.Parent.Data != "style" {
			fmt.Println(top.Data)
		}
	}
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
