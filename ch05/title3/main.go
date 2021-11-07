package main

import (
	"fmt"

	"golang.org/x/net/html"
)

// soleTitleはdoc中の最初の空ではないtitle要素のテキストと、
// title要素が一つだけでなかったエラーを返します
func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
		// パニック無し
		case bailout{}:
			// 「予期された」パニック
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p) // 予期しないパニック; パニックを維持する
		}
	}()

	// 二つ以上の空ではないtitleを見つけたら再起から抜け出させる。
	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" &&
			n.FirstChild != nil {
			if title != "" {
				panic(bailout{}) // 複数のtitle要素
			}
			title = n.FirstChild.Data
		}
	}, nil)
	if title == "" {
		return "", fmt.Errorf("no title element")
	}
	return title, nil
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
