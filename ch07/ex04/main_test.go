package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestNewReader(t *testing.T) {
	tests := []string{
		"",
		"foo",
		"foo bar",
		"日本語",
		strings.Repeat("A", 10000),
	}

	for _, test := range tests {
		r := NewReader(test)
		var b bytes.Buffer
		n, err := io.Copy(&b, r) // bにコピーできれば自作のReaderがio.Readerを実装できていることを確認できる
		if err != nil {
			t.Error(err)
			continue
		}
		if got := int(n); got != len(test) {
			t.Errorf("unexpected write bytes. expected: %v, but got: %v", len(test), got)
		}
		if got := b.String(); got != test {
			t.Errorf("unexpected result. expected: %q, but got: %q", test, got)
		}
	}
}
