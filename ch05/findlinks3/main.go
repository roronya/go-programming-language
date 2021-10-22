package main

import (
	"fmt"
	"log"
	"os"

	"github.com/roronya/go-programming-language/ch05/links"
)

// breadthFirstはworklist内の個々の項目に対してfを呼び出します
// fから返されたすべての項目はworklistへ追加されます
// fはそれぞれの項目に対して高々一度しか呼び出されません。
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}
