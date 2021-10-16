package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestComma_3桁以下のとき(t *testing.T) {
	actual := comma(*bytes.NewBufferString("123"))
	if actual != "123" {
		t.Errorf("%s is not 123", actual)
	}
}

func TestComma_3桁以上のとき(t *testing.T) {
	actual := comma(*bytes.NewBufferString("1234567"))
	if actual != "1,234,567" {
		t.Errorf("%s is not 1,234,567", actual)
	}
}

func TestComma_小数点があるとき(t *testing.T) {
	actual := comma(*bytes.NewBufferString("1234.56789"))
	if actual != "1,234.56789" {
		t.Errorf("%s is not 1,234.56789", actual)
		strings.Count()
	}
}
