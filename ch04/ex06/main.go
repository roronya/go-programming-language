package ex06

import "unicode"

func f(slice []byte) []byte {
	i := 0
	j := 0
	for i < len(slice) {
		if unicode.IsSpace(rune(slice[i])) {
			slice[j] = ' '
			i++
			j++
			for i < len(slice) && unicode.IsSpace(rune(slice[i])) {
				i++
			}
		}

	}
	return slice[:j]
}
