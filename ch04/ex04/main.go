package main

func rotate(a []int) []int {
	return append(a, a[0])[1:]
}
