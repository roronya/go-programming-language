package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestCountWordsAndImages(t *testing.T) {
	f, _ := os.Open("test.html")
	txt, _ := ioutil.ReadAll(f)
	doc, _ := html.Parse(bytes.NewBuffer(txt))

	words, images := countWordsAndImages(doc)
	if words != 14 || images != 2 {
		log.Fatalf("error! words: %d, images: %d", words, images)
	}
}
