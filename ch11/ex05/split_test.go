package split

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		s    string
		sep  string
		want int
	}{
		{"a:b:c", ":", 3},
		// ...
	}
	for _, test := range tests {
		if got := strings.Split(test.s, test.sep); len(got) != test.want {
			t.Errorf("Split(%q, %q) returned %d words, want %d",
				test.s, test.sep, len(got), test.want)
		}
	}
}
