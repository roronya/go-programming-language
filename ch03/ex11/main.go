package main

import (
	"bytes"
	"fmt"
)

func main() {
	values := *bytes.NewBufferString("1234567")
	result := comma(values)
	fmt.Printf(result)
}

func comma(values bytes.Buffer) string {
	var intPart []byte
	var decimalPart []byte
	dotIndex := bytes.IndexByte(values.Bytes(), '.')
	if dotIndex < 0 {
		intPart = values.Bytes()[:values.Len()]
	} else {
		intPart = values.Bytes()[:dotIndex]
		decimalPart = values.Bytes()[dotIndex:]
	}
	var buf bytes.Buffer
	if len(intPart) <= 3 {
		buf.Write(intPart)
		buf.Write(decimalPart)
		return buf.String()
	}
	for i, v := range intPart {
		j := values.Len() - i
		if j%3 == 0 {
			buf.WriteString(",")
		}
		buf.WriteByte(v)
	}
	buf.Write(decimalPart)
	return buf.String()
}
