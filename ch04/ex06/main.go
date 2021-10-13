package ex06

import "unicode"

func f(slice []byte) []byte {
	for i, v := range slice {
		if unicode.IsSpace(rune(v)) {
			slice[i] = ' '
		}
	}
	return slice
}
