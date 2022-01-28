package display

import (
	"os"
	"reflect"
)

func ExampleMovie() {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Color           bool
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}

	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},

		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	Display("strangelove", strangelove)
	//Output:
}

func ExampleFile() {
	Display("os.Stderr", os.Stderr)
	//Output:
}

func ExampleRV() {
	Display("rV", reflect.ValueOf(os.Stderr))
	//Output:
}

func ExampleInterface() {
	var i interface{} = 3

	Display("i", i)
	//Output:
	//Display i (int):
	//i = 3
}

func ExampleInterfacePointer() {
	var i interface{} = 3

	Display("&i", &i)
	//Output:
	//Display &i (*interface {}):
	//(*&i).type = int
	//(*&i).value = 3
}

// Never end
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
