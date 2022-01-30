package display

func ExampleDisplay() {
	type Movie struct {
		Title string
		Year  int
	}
	m := map[Movie]int{
		Movie{"a", 2001}: 1,
		Movie{"b", 2002}: 2,
		Movie{"c", 2003}: 3,
	}

	Display("m", m)
	//Output:
}
