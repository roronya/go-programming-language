package format

import (
	"fmt"
	"time"

	"gopl.io/ch12/format"
)

func ExampleAny() {
	var x int64 = 1
	var d time.Duration = 1 * time.Nanosecond
	fmt.Println(format.Any(x))
	fmt.Println(format.Any(d))
	fmt.Println(format.Any([]int64{}))
	fmt.Println(format.Any([]time.Duration{d}))
	//Output:
}
