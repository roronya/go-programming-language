package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		log.Fatalf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatalf("parsing %s as HTML: %v", url, err)
	}
	forEachNode(doc, startElement, endElement)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int
var out io.Writer = os.Stdout

func startElement(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		printIndent(depth)
		fmt.Fprintf(out, "<%s", n.Data)
		for _, a := range n.Attr {
			fmt.Fprintf(out, " %s=\"%s\"", a.Key, a.Val)
		}
		if c := n.FirstChild; c == nil {
			fmt.Println("/>")
			return
		}
		fmt.Println(">")
		depth++
	case html.TextNode:
		printIndent(depth)
		fmt.Fprintf(out, "%s\n", n.Data)
	case html.CommentNode:
		printIndent(depth)
		fmt.Fprintf(out, "<!--%s-->\n", n.Data)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		if c := n.FirstChild; c != nil {
			depth--
			printIndent(depth)
			fmt.Fprintf(out, "</%s>\n", n.Data)
		}
	}
}

func printIndent(depth int) {
	fmt.Fprintf(out, "%*s", depth*2, "")
}
