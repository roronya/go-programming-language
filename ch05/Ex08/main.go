package main

import (
	"golang.org/x/net/html"
)

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	if pre != nil && pre(n) {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if forEachNode(c, pre, post) != nil {
			return c
		}
	}

	if post != nil && post(n) {
		return n
	}
	return nil
}

func ElementByID(doc *html.Node, id string) *html.Node {
	pre := func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == id {
					return true
				}
			}
		}
		return false
	}
	return forEachNode(doc, pre, nil)
}
