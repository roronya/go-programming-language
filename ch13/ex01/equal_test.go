package equal

import "testing"

func TestEqual(t *testing.T) {
	tests := []struct {
		x, y interface {
		}
		want bool
	}{
		{1, 1, true},
		{0, 1, false},
		{1.0, 1.0, true},
		{1.0, 1.1, false},
		{1e-9, 1e-9, true},
		{1e-9, 1.1e-9, true},
		{"abc", "abc", false},
		{1e-9, "abc", false},
	}
	for _, test := range tests {
		if got := Equal(test.x, test.y); got != test.want {
			t.Errorf("case x:%#v y:%#v - want %#v, got %#v", test.x, test.y, test.want, got)
		}
	}
}
