package main

import (
	"log"
	"testing"
)

// strings_test.goと同じケースをテストする
var abcd = "abcd"
var faces = "☺☻☹"
var commas = "1,2,3,4"
var dots = "1....2....3....4"

var tests = []struct {
	expected string
	sep      string
	in       []string
}{
	{"", "", []string{}},
	{abcd, "", []string{"a", "bcd"}},
	{abcd, "", []string{"a", "b", "c", "d"}},
	{abcd, "", []string{"a", "b", "c", "d"}},
	{faces, "", []string{"☺", "☻", "☹"}},
	{faces, "", []string{"☺", "☻", "☹"}},
	{faces, "", []string{"☺", "☻", "☹"}},
	{"☺�☹", "", []string{"☺", "�", "☹"}},
	{"", "a", nil},
	{abcd, "a", []string{"", "bcd"}},
	{abcd, "z", []string{"abcd"}},
	{commas, ",", []string{"1", "2", "3", "4"}},
	{dots, "...", []string{"1", ".2", ".3", ".4"}},
	{faces, "☹", []string{"☺☻", ""}},
	{faces, "~", []string{faces}},
	{"1 2 3 4", " ", []string{"1", "2", "3 4"}},
	{"1 2", " ", []string{"1", "2"}},
}

func TestJoin(t *testing.T) {
	for i, test := range tests {
		actual := join(test.sep, test.in...)
		if actual != test.expected {
			log.Fatalf("failed case%d: expect %s but got %s\n", i, test.expected, actual)
		}
	}
}
