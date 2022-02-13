package search

import "testing"

func TestPack(t *testing.T) {
	in := struct {
		X int        `http:"X"`
		Y string     `http:"Y"`
		Z complex128 `ftp:"Z"`
	}{
		1000,
		"abc",
		complex(1, 1),
	}
	got := Pack(&in)
	want := "?X=1000&Y=abc"
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}
