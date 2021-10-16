package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"testing"

	"golang.org/x/net/html"
)

func TestVisit(t *testing.T) {
	f, err := os.Open("test.html")
	if err != nil {
		log.Fatal(err)
	}

	txt, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := html.Parse(bytes.NewBuffer(txt))
	if err != nil {
		fmt.Fprintf(os.Stderr, "html parse failed: %v\n", err)
		os.Exit(1)
	}

	actual := visit(doc)
	expected := []string{"link0", "link1", "link2", "link3", "link4", "link5"}
	sort.Strings(actual)
	sort.Strings(expected)
	for i, _ := range actual {
		if actual[i] != expected[i] {
			log.Fatalf("%s != %s", actual[i], expected[i])
		}
	}

}
