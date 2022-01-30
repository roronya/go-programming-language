package display

// 3層まで表示して止まる
func ExampleCycle() {
	type Cycle struct {
		Value int
		Tail  *Cycle
	}
	var c Cycle
	c = Cycle{42, &c}
	Display("c", c)
	//Output:
}
