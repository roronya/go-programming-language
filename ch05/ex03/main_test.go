package main

import (
	"log"
	"testing"

	"golang.org/x/net/html"
)

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
