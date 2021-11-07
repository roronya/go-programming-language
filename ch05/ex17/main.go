package main

import "golang.org/x/net/html"

func ElementByTagName(doc *html.Node, names ...string) []*html.Node {
	var result, q []*html.Node
	q = append(q, doc)
	for len(q) > 0 {
		worklist := q
		q = nil
		for _, node := range worklist {
			for c := node.FirstChild; c != nil; c = c.NextSibling {
				q = append(q, c)
			}

			if node.Type == html.ElementNode {
				for _, name := range names {
					if name == node.Data {
						result = append(result, node)
					}
				}
			}
		}
	}
	return result
}
