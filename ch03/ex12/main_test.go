package main

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	var actual bool
	var a, b string

	a = "しんぶんし"
	b = "しんぶんし"
	actual = isAnagram(a, b)
	if !actual {
		t.Errorf("%sと%sがアナグラムではないと判定された", a, b)
	}

	a = "アナグラム"
	b = "グアムナラ"
	actual = isAnagram(a, b)
	if !actual {
		t.Errorf("%sと%sがアナグラムではないと判定された", a, b)
	}
	a = "あいうえお"
	b = "あいうえあ"
	actual = isAnagram(a, b)
	if actual {
		t.Errorf("%sと%sがアナグラムだと判定された", a, b)
	}

}
