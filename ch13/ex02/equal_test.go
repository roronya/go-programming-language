package equal

import "testing"

func TestIsCyclic(t *testing.T) {
	type link struct {
		value string
		tail  *link
	}

	a, b, c, d, e := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}, &link{value: "d"}, &link{value: "e"}
	a.tail, b.tail, c.tail, d.tail = b, a, c, e

	tests := []struct {
		in   *link
		want bool
	}{
		{a, true},  // a -> b -> a -> ...
		{b, true},  // b -> a -> b -> ...
		{c, true},  // c -> c -> c -> ...
		{d, false}, // d -> e -> nil
	}
	for _, test := range tests {
		if got := IsCyclic(test.in); got != test.want {
			t.Errorf("case %#v: want %#v, got %#v", test.in, test.want, got)
		}
	}
}
