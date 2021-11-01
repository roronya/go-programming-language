package main

import (
	"io"
	"log"
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
			"./dir0/dir1/test.txt",
			[]byte{'e', 'x', 'p', 'e', 'c', 't', 'e', 'd'},
			"dir0",
		},
		{
			"./test.txt",
			[]byte{'e', 'x', 'p', 'e', 'c', 't', 'e', 'd'},
			"test.txt",
		},
		{
			"./dir0/../test.txt",
			[]byte{'e', 'x', 'p', 'e', 'c', 't', 'e', 'd'},
			"test.txt",
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

func TestDownload(t *testing.T) {
	err := download("https://example.com")
	if err != nil {
		log.Fatal(err)
	}
}
