package sexpr

import (
	"encoding/json"
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"Hello, World", `"Hello, World"`},
		{"", `""`},
		{`Hello, "World"`, `"Hello, \"World\""`},
		{`Hello, \World\`, `"Hello, \\World\\"`},
	}
	for _, test := range tests {
		b, err := Marshal(test.in)
		if err != nil {
			t.Errorf("err! %s: %s\n", test.in, err)
		}
		if string(b) != test.want {
			t.Errorf("want %s, b %s\n", test.want, b)
		}
		var actual string
		err = json.Unmarshal(b, &actual)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestNumber(t *testing.T) {
	tests := []struct {
		in   interface{}
		want string
	}{
		{42, `42`},
		{-42, `-42`},
		{1.000, "1"},
		{6.02214129e23, "6.02214129e+23"},
		{6.62606957e-34, "6.62606957e-34"},
	}
	for _, test := range tests {
		b, err := Marshal(test.in)
		if err != nil {
			t.Errorf("err! %s: %s\n", test.in, err)
		}
		if string(b) != test.want {
			t.Errorf("want %s, b %s\n", test.want, b)
		}
		var actual interface{}
		err = json.Unmarshal(b, &actual)
		if err != nil {
			t.Errorf("cannot unmarshal  in:%v got:%s err:%s\n", test.in, string(b), err)
		}
	}
}

func TestBool(t *testing.T) {
	tests := []struct {
		in   bool
		want string
	}{
		{true, `true`},
		{false, `false`},
	}
	for _, test := range tests {
		b, err := Marshal(test.in)
		if err != nil {
			t.Errorf("err! %t: %s\n", test.in, err)
		}
		if string(b) != test.want {
			t.Errorf("want %s, b %s\n", test.want, b)
		}
		var actual interface{}
		err = json.Unmarshal(b, &actual)
		if err != nil {
			t.Errorf("cannot unmarshal  in:%v got:%s err:%s\n", test.in, string(b), err)
		}
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		in   interface{}
		want string
	}{
		{map[string]string{"key": "value"}, `{ "key": "value" }`},
		// ↓キーの順番は保証できないから確率的に落ちる
		{map[string]string{"key1": "value1", "key2": "value2"}, `{ "key1": "value1", "key2": "value2" }`},
		{map[string]map[string]string{"key": {"key": "value"}}, `{ "key": { "key": "value" } }`},
		{map[string]string{}, `{ }`},
	}
	for _, test := range tests {
		b, err := Marshal(test.in)
		if err != nil {
			t.Errorf("err! %s: %s\n", test.in, err)
		}
		if string(b) != test.want {
			t.Errorf("want %s, b %s\n", test.want, b)
		}
		var actual interface{}
		err = json.Unmarshal(b, &actual)
		if err != nil {
			t.Errorf("cannot unmarshal  in:%v b:%s err:%s\n", test.in, string(b), err)
		}
	}
}

func TestObject(t *testing.T) {
	tests := []struct {
		in   interface{}
		want string
	}{
		{struct{ key string }{"value"}, `{ "key": "value" }`},
		{struct {
			key1 string
			key2 string
		}{"value1", "value2"}, `{ "key1": "value1", "key2": "value2" }`},
		{struct {
			key struct {
				key string
			}
		}{struct{ key string }{"value"}}, `{ "key": { "key": "value" } }`},
		{struct{}{}, `{ }`},
	}
	for _, test := range tests {
		b, err := Marshal(test.in)
		if err != nil {
			t.Errorf("err! %s: %s\n", test.in, err)
		}
		if string(b) != test.want {
			t.Errorf("want %s, b %s\n", test.want, b)
		}
		var actual interface{}
		err = json.Unmarshal(b, &actual)
		if err != nil {
			t.Errorf("cannot unmarshal  in:%v b:%s err:%s\n", test.in, string(b), err)
		}
	}
}
func TestStrangelove(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
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

	actual := Movie{}

	b, err := Marshal(strangelove)
	t.Logf("marshal: \n%s\n", b)
	if err != nil {
		t.Error(err)
	}
	err = json.Unmarshal(b, &actual)
	if err != nil {
		t.Error(err)
	}
	if actual.Title != strangelove.Title {
		t.Errorf("error")
	}
}
