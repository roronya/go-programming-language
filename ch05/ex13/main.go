package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

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
	domain := getDomain(url)
	// FIXME: links.Extractと合わせて二度リクエストを投げてしまっている
	err := download(url)

	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}

	// 元のurlと違うドメインのurlがlistに入っていればそれをフィルタしておけばよい
	list = filterByPrefix(list, domain)

	return list
}

func download(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	// TODO: コンテンツの保存をする
	err = save(resp.Body, url)
	if err != nil {
		return err
	}
	return nil
}

func save() error {
	/**
	err = os.Mkdir(domain) // FIXME: 既にあるディレクトリを作らないようにする
	if err != nil {
		log.Print(err)
	}
	*/

	// pathごとにディレクトリを切ってコンテンツを保存する
	return nil
}

func filterByPrefix(list []string, prefix string) []string {
	result := []string{}
	for _, l := range list {
		if !strings.HasPrefix(l, prefix) {
			continue
		}
		result = append(result, l)
	}
	return result
}

func getDomain(url string) string {
	return "gopl.io"
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}
