package main

import (
	"fmt"
	"os"
	"strconv"
)

type Centi float64
type Inch float64

func CtoI(c Centi) Inch {
	return Inch(c * 0.39370)
}

func ItoC(i Inch) Centi {
	return Centi(i / 0.39370)
}

func (c Centi) String() string {
	return fmt.Sprintf("%g cm", c)
}

func (i Inch) String() string {
	return fmt.Sprintf("%g in", i)
}

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		c := Centi(t)
		i := Inch(t)
		fmt.Printf("%s = %s, %s = %s\n",
			c, CtoI(c), i, ItoC(i))
	}
}
