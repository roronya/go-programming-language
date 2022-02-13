package sexpr

import (
	"bytes"
	"testing"
)

func TestToken(t *testing.T) {
	//in := []byte(`((Title "Dr. Strangelove") (Subtitle "How I Learned to Stop Worrying and Love the Bomb") (Year 1964) (Actor (("Dr. Strangelove" "Peter Sellers") ("Grp. Capt. Lionel Mandrake" "Peter Sellers") ("Pres. Merkin Muffley" "Peter Sellers") ("Gen. Buck Turgidson" "George C. Scott") ("Brig. Gen. Jack D. Ripper" "Sterling Hayden") ("Maj. T.J. \"King\" Kong" "Slim Pickens"))) (Oscars ("Best Actor (Nomin.)" "Best Adapted Screenplay (Nomin.)" "Best Director (Nomin.)" "Best Picture (Nomin.)")) (Sequel nil))`)
	in := []byte(`((X "x") (Y (1 2 3)))'`)
	want := [...]Token{
		StartList('('),
		StartList('('),
		Symbol("X"),
		String("x"),
		EndList(')'),
		StartList('('),
		Symbol("Y"),
		StartList('('),
		Int(1),
		Int(2),
		Int(3),
		EndList(')'),
		EndList(')'),
		EndList(')'),
	}
	dec := NewDecoder(bytes.NewReader(in))
	var got []Token
	for i := 0; i < len(want); i++ {
		token, err := dec.Token()
		if err != nil {
			t.Error(err)

		}
		got = append(got, token)
	}
	for i := 0; i < len(want); i++ {
		if want[i] != got[i] {
			t.Logf("want %#v, got %#v", want, got)
		}
	}
}
