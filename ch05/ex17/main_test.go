package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestElementByTagName(t *testing.T) {
	f, _ := os.Open("test.html")
	txt, _ := ioutil.ReadAll(f)
	doc, _ := html.Parse(bytes.NewBuffer(txt))
	imgs := ElementByTagName(doc, "img")
	if len(imgs) != 10 {
		log.Fatalf("img tag size is expected 10, but got %d\n", len(imgs))
	}
	for _, img := range imgs {
		if img.Data != "img" {
			log.Fatalf("expected img tag, but got %s\n", img.Data)
		}
	}

	headers := ElementByTagName(doc, "h1", "h2", "h3", "h4")
	if len(headers) != 10 {
		log.Fatalf("img h tag size is expected 10, but got %d\n", len(headers))
	}
	for _, h := range headers {
		if h.Data != "h1" && h.Data != "h2" && h.Data != "h3" && h.Data != "h4" {
			log.Fatalf("expected h tag, but got %s\n", h.Data)
		}
	}
}
