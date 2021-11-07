package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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

func crawl(rawUrl string) []string {
	err := download(rawUrl)
	if err != nil {
		log.Print(err)
	}

	list, err := links.Extract(rawUrl)
	if err != nil {
		log.Print(err)
	}

	// 元のurlと違うドメインのurlがlistに入っていればそれをフィルタして、異なるドメインを探索しないようにする
	url_, err := url.Parse(rawUrl)
	if err != nil {
		log.Print(err)
	}
	list = selectByPrefix(list, url_.Scheme+"://"+url_.Hostname())

	for _, item := range list {
		fmt.Println(item)
	}

	return list
}

func download(url_ string) error {
	resp, err := http.Get(url_)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("getting %s: %s", url_, resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	path := urlToFilepath(resp.Request.URL)
	log.Println(path)
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

func selectByPrefix(list []string, prefix string) []string {
	result := []string{}
	for _, l := range list {
		if !strings.HasPrefix(l, prefix) {
			continue
		}
		result = append(result, l)
	}
	return result
}

func urlToFilepath(u *url.URL) string {
	builder := strings.Builder{}
	builder.WriteString(u.Hostname())
	builder.WriteString(u.Path)
	// 末尾が省略されてるケースはindex.htmlとして処理する
	// pathが指定されないケースケース 例: https://www.debian.or.jp
	if u.Path == "" {
		builder.WriteString("/index.html")
	}
	// 末尾が/で終わるケース 例: https://www.debian.or.jp/doc/
	if strings.HasSuffix(u.Path, "/") {
		builder.WriteString("index.html")
	}
	return builder.String()
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}
