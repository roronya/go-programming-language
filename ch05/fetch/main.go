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
	// f.Close()でエラーを報告される可能性があるので、
	// defer f.Close()のようにしてしまうとエラーを握りつぶしつつ、Close()にも失敗してしまう

	n, err = io.Copy(f, resp.Body)
	// ファイルを閉じるが、Copyでエラーがあればそちらを優先する
	// ファイルを明示的に閉じつつ、エラーが報告したいが、io.Copyのほうが有益なエラーの可能性があるので優先する
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}
