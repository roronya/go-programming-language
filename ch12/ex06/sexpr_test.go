package sexpr

import (
	"testing"
)

func TestMarshal(t *testing.T) {
	tests := []struct {
		x    interface{}
		want string
	}{
		{struct{ X, Y int }{Y: 1}, "((Y 1))"},
		{struct{ X, Y string }{Y: "ab"}, "((Y \"ab\"))"},
		{struct{ X, Y byte }{Y: 'a'}, "((Y 97))"},
		{struct{ X, Y []int }{Y: []int{1, 2}}, "((Y (1 2)))"},
		{struct{ X, Y struct{ X, Y int } }{Y: struct{ X, Y int }{Y: 1}}, "((Y ((Y 1))))"},
	}
	for _, test := range tests {
		actual, _ := Marshal(test.x)
		if string(actual) != test.want {
			t.Errorf("want %s, got %s", test.want, actual)
		}
	}
}
