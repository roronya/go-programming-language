package main

import (
	"io"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"testing"
)

// saveで指定したディレクトリにファイルを作られるかテストする
func TestSave(t *testing.T) {
	tests := []struct {
		path     string
		expected []byte
		remove   string
	}{
		{
			"www.example.com/dir0/dir1/test.txt",
			[]byte{'e', 'x', 'p', 'e', 'c', 't', 'e', 'd'},
			"www.example.com",
		},
		{
			"www.example.com/test.txt",
			[]byte{'e', 'x', 'p', 'e', 'c', 't', 'e', 'd'},
			"www.example.com",
		},
		{
			"www.example.com/dir0/../test.txt",
			[]byte{'e', 'x', 'p', 'e', 'c', 't', 'e', 'd'},
			"www.example.com",
		},
	}

	for _, test := range tests {
		err := save(test.expected, test.path)
		if err != nil {
			log.Fatal(err)
		}

		f, err := os.Open(filepath.Clean(test.path))
		if err != nil {
			log.Fatal(err)
		}
		actual, err := io.ReadAll(f)
		if err != nil {
			log.Fatal(err)
		}

		os.RemoveAll(test.remove)

		if len(test.expected) != len(actual) {
			log.Fatalf("testcase %s: expected %d. but got %d\n", test.path, len(test.expected), len(actual))
		}
		for i, _ := range actual {
			if actual[i] != test.expected[i] {
				log.Fatalf("testcase %s: expected %b. but got %b\n", test.path, test.expected[i], actual[i])
			}
		}
	}
}

// ダウンロードできるか
func TestDownload(t *testing.T) {
	err := download("https://example.com")
	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Stat("example.com/index.html")
	if err != nil {
		log.Fatal("download is failed")
	}

	os.RemoveAll("example.com")
}

func TestFilterByPrefix(t *testing.T) {
	tests := []struct {
		list     []string
		prefix   string
		expected []string
	}{
		// よくありそうなケース
		{
			[]string{"https://a.com/index.html", "https://b.org/doc/index.html", "https://a.com/doc/index.html"},
			"https://a.com",
			[]string{"https://a.com/index.html", "https://a.com/doc/index.html"},
		},
		// ポート指定する
		{
			[]string{"https://a.com:8080/index.html", "https://b.org/doc/index.html", "https://a.com:8080/doc/index.html"},
			"https://a.com:8080",
			[]string{"https://a.com:8080/index.html", "https://a.com:8080/doc/index.html"},
		},
		// prefixを持つURLがlistの中に存在しないケース
		{
			[]string{"https://a.com/index.html", "https://b.org/doc/index.html", "https://a.com/doc/index.html"},
			"https://z.com",
			[]string{},
		},
	}

	for i, test := range tests {
		actual := selectByPrefix(test.list, test.prefix)
		if len(test.expected) != len(actual) {
			log.Fatalf("testcase %d: expected %d, but got %d", i, len(test.expected), len(actual))
		}
		for j, _ := range test.expected {
			if test.expected[j] != actual[j] {
				log.Fatalf("testcase %d: expected %s, but got %s", i, test.expected, actual)
			}
		}
	}

}

func TestUrlToFilepath(t *testing.T) {
	tests := []struct {
		title    string
		in       string
		expected string
	}{
		{
			"末尾にパスが無いケース",
			"https://example.com",
			"example.com/index.html",
		},
		{
			"末尾がスラッシュで終わるケース",
			"https://example.com/",
			"example.com/index.html",
		},
		{
			"末尾にパスがあるケース",
			"https://example.com/index.html",
			"example.com/index.html",
		},
		{
			"末尾のパスがindex.htmlではないケース",
			"https://example.com/help.html",
			"example.com/help.html",
		},
		{
			"パスの階層が深いケース",
			"https://example.com/a/b/c/d/e/f/g/h/",
			"example.com/a/b/c/d/e/f/g/h/index.html",
		},
	}

	for _, test := range tests {
		u, _ := url.Parse(test.in) // ignore err
		actual := urlToFilepath(u)
		if actual != test.expected {
			log.Fatalf("expected %s, but got %s\n", test.expected, actual)
		}
	}
}

func TestCrawl(m *testing.T) {
	breadthFirst(crawl, []string{"https://example.com"})
	_, err := os.Stat("example.com/index.html")
	if err != nil {
		log.Fatal("crawl is failed")
	}
	os.RemoveAll("example.com")
}
