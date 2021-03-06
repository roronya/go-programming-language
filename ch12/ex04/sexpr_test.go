package sexpr

import (
	"fmt"
	"log"
)

func ExampleMarshal() {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}

	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
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

	b, err := Marshal(strangelove)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", b)
	//Output:
	//((Title "Dr. Strangelove")
	// (Subtitle "How I Learned to Stop Worrying and Love the Bomb")
	// (Year 1964)
	// (Actor (("Gen. Buck Turgidson" "George C. Scott")
	//         ("Brig. Gen. Jack D. Ripper" "Sterling Hayden")
	//         ("Maj. T.J. \"King\" Kong" "Slim Pickens")
	//         ("Dr. Strangelove" "Peter Sellers")
	//         ("Grp. Capt. Lionel Mandrake" "Peter Sellers")
	//         ("Pres. Merkin Muffley" "Peter Sellers")))
	// (Oscars ("Best Actor (Nomin.)"
	//          "Best Adapted Screenplay (Nomin.)"
	//          "Best Director (Nomin.)"
	//          "Best Picture (Nomin.)"))
	// (Sequel nil))
}
