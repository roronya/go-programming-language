package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

type Reader struct {
	s string
	i int
}

func (r *Reader) Read(p []byte) (n int, err error) {
	n = copy(p, r.s[r.i:])
	r.i += n
	if n != len(p) {
		return n, io.EOF
	}
	return n, nil
}

func NewReader(s string) io.Reader {
	return &Reader{s: s}
}

func main() {
	for _, in := range os.Args[1:] {
		doc, err := html.Parse(NewReader(in))
		if err != nil {
			fmt.Fprintf(os.Stdin, "findlinnks1: %v\n", err)
			os.Exit(1)
		}
		for _, link := range visit(nil, doc) {
			fmt.Println(link)
		}
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
