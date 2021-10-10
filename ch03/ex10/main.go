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
	if values.Len() <= 3 {
		return values.String()
	}
	var buf bytes.Buffer
	for i, v := range values.Bytes() {
		j := values.Len() - i
		if j%3 == 0 {
			buf.WriteString(",")
		}
		buf.WriteByte(v)
	}
	return buf.String()
}
