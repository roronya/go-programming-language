package sexpr

import (
	"reflect"
	"testing"
)

func TestTag(t *testing.T) {
	type Movie struct {
		Title    string `sexpr:"title"`
		Subtitle string `sexpr:"subtitle"`
		Year     int
		Actor    map[string]string
		Oscars   []string
		Sequel   *string
	}

	in := Movie{
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

	b, err := Marshal(in)
	if err != nil {
		t.Error(err)
	}

	got := Movie{}
	err = Unmarshal(b, &got)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(in, got) {
		t.Errorf("want %#v, got %#v", in, got)
	}
}
