package main

import "golang.org/x/net/html"

func ElementByTagName(doc *html.Node, name ...string) []*html.Node {
	var result []*html.Node
	q := []*html.Node{doc}
	var top *html.Node
	for !isEmpty(q) {
		q, top = pop(q)
		for c := top.FirstChild; c != nil; c = c.NextSibling {
			q = push(q, c)
		}

		if top.Type == html.ElementNode {
			for _, n := range name {
				if n == top.Data {
					result = append(result, top)
				}
			}
		}
	}
	return result
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
