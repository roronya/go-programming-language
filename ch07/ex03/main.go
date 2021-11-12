package main

import (
	"strconv"
	"strings"
)

type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() string {
	var builder strings.Builder
	builder.WriteString("[")
	builder.WriteString(strconv.Itoa(t.value))

	builder.WriteString(" ")
	if t.left != nil {
		builder.WriteString(t.left.String())
	} else {
		builder.WriteString("[]")
	}

	builder.WriteString(" ")
	if t.right != nil {
		builder.WriteString(t.right.String())
	} else {
		builder.WriteString("[]")
	}
	builder.WriteString("]")
	return builder.String()
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}
