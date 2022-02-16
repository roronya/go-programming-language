package equal

import "testing"

func TestEqual(t *testing.T) {
	tests := []struct {
		x, y interface{}
		want bool
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}, true},
		{[]string{"foo"}, []string{"bar"}, false},
		{[]string(nil), []string{}, true},
		{map[string]int(nil), map[string]int{}, true},
	}
	for _, test := range tests {
		if got := Equal(test.x, test.y); got != test.want {
			t.Errorf("want %#v, got %#v", test.want, got)
		}
	}
}

func TestEqualCyclicStruct(t *testing.T) {
	type link struct {
		value string
		tail  *link
	}

	a, b, c := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}
	a.tail, b.tail, c.tail = b, a, c

	tests := []struct {
		x, y *link
		want bool
	}{
		{a, a, true},
		{b, b, true},
		{c, c, true},
		{a, b, false},
		{a, c, false},
	}
	for _, test := range tests {
		if got := Equal(test.x, test.y); got != test.want {
			t.Errorf("want %#v, got %#v", test.want, got)
		}
	}
}
