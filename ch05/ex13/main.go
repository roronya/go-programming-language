package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/roronya/go-programming-language/ch05/links"
)

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
	// FIXME: links.Extractと合わせて二度リクエストを投げてしまっている
	err := download(url)

	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}

	// 元のurlと違うドメインのurlがlistに入っていればそれをフィルタしておけばよい
	domain := getDomain(url)
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	// TODO: コンテンツの保存をする
	path := "./" + resp.Request.URL.Path
	err = save(body, path)
	if err != nil {
		return err
	}
	return nil
}

func save(bytes []byte, path string) error {
	path = filepath.Clean(path)
	if err := os.MkdirAll(filepath.Dir(path), 0777); err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	if _, err := f.Write(bytes); err != nil {
		return err
	}

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
