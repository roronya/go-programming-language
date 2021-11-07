package main

import (
	"io"
	"net/http"
	"os"
	"path"
)

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	// f.Close()しつつもio.Copy()でエラーがあったら優先するような無名関数を作ってdeferにわたす
	fileClose := func() {
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}
	defer fileClose()

	n, err = io.Copy(f, resp.Body)
	return local, n, err
}
