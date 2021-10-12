package main

import (
	"sort"
	"strings"
)

func isAnagram(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}
	s := strings.Split(a, "")
	t := strings.Split(b, "")
	sort.Strings(s)
	sort.Strings(t)
	return strings.Join(s, "") == strings.Join(t, "")
}
