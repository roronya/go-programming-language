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
	expected := map[string]int{"title": 1, "meta": 3, "style": 1, "head": 1, "body": 1, "div": 1, "h1": 1, "p": 2, "a": 1}

	var keys []string
	for key := range expected {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		if actual[key] != expected[key] {
			log.Fatalf("key %s: %d != %d", key, actual[key], expected[key])
		}
	}

}
func TestPush(t *testing.T) {
	e1 := &html.Node{}
	e2 := &html.Node{}
	var q []*html.Node
	q = push(q, e1)
	q = push(q, e2)
	if len(q) != 2 {
		log.Fatal("error")
	}
	expected := []*html.Node{e1, e2}
	for i, _ := range q {
		if q[i] != expected[i] {
			log.Fatalf("%v != %v", q[i], expected[i])
		}
	}
}
func TestPop(t *testing.T) {
	e1 := &html.Node{}
	e2 := &html.Node{}
	q := []*html.Node{e1, e2}
	q, actual := pop(q)
	if actual != e1 {
		log.Fatal("error")
	}
	if len(q) != 1 {
		log.Fatal("error")
	}
	q, actual = pop(q)
	if actual != e2 {
		log.Fatal("error")
	}
	if len(q) != 0 {
		log.Fatal("error")
	}
}
