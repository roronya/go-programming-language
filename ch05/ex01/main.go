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
	for _, link := range visit(doc) {
		fmt.Println(link)
	}
}

func visit(n *html.Node) (l []string) {
	// 末尾まで来たら何もせず返す
	if n == nil {
		return
	}
	// 子ノードを探索する
	l = append(l, visit(n.FirstChild)...)
	// 兄弟ノードを探索する
	l = append(l, visit(n.NextSibling)...)

	// 今のノードを調べてリンクがあれば返り値に追加する
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				l = append(l, a.Val)
			}
		}
	}
	return
}
