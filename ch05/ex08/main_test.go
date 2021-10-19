package main

import (
	"bytes"
	"log"
	"testing"

	"golang.org/x/net/html"
)

func TestElementByID(t *testing.T) {
	doc, _ := html.Parse(bytes.NewBufferString("<p id=\"target\">this is test</p>"))
	if actual := ElementByID(doc, "target"); actual == nil {
		log.Fatalf("target is not found. got=%#v", actual)
	}
	if actual := ElementByID(doc, "notExistId"); actual != nil {
		log.Fatalf("notExistId is found. got=%#v", actual)
	}
}
