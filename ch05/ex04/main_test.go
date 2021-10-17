package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"testing"

	"golang.org/x/net/html"
)

func TestVisit(t *testing.T) {
	f, _ := os.Open("test.html")
	txt, _ := ioutil.ReadAll(f)
	doc, _ := html.Parse(bytes.NewBuffer(txt))

	actual := visit(doc)
	expected := []string{"link0", "link1", "link2", "link3", "link4", "link5", "link6", "link7", "link8"}
	sort.Strings(actual)
	sort.Strings(expected)
	for i, _ := range actual {
		if actual[i] != expected[i] {
			log.Fatalf("%s != %s\nactual: %#v\nexpected: %#v", actual[i], expected[i], actual, expected)
		}
	}

}
