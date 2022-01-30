package sexpr

import (
	"log"
	"testing"
)

func TestBool(t *testing.T) {
	actual, _ := Marshal(true)
	if string(actual) != "t" {
		log.Fatalf("want \"t\", got \"%s\"", actual)
	}

	actual, _ = Marshal(false)
	if string(actual) != "nil" {
		log.Fatalf("want \"nil\", got \"%s\"", actual)
	}
}

func TestFloat64(t *testing.T) {
	tests := []struct {
		in   float64
		want string
	}{
		{1.000, "1"},
		{6.02214129e23, "6.02214129e+23"},
		{6.62606957e-34, "6.62606957e-34"},
	}
	for _, test := range tests {
		actual, _ := Marshal(test.in)
		if string(actual) != test.want {
			log.Fatalf("want \"%s\", got \"%s\"", test.want, actual)
		}
	}
}

func TestComplex(t *testing.T) {
	tests := []struct {
		in   complex128
		want string
	}{
		{complex(1, 2), "#(1 2)"},
		{complex(1.1, 2.1), "#(1.1 2.1)"},
		{complex(0, 0), "#(0 0)"},
		{complex(1, 0), "#(1 0)"},
		{complex(0, 1), "#(0 1)"},
	}
	for _, test := range tests {
		actual, _ := Marshal(test.in)
		if string(actual) != test.want {
			log.Fatalf("want \"%s\", got \"%s\"", test.want, actual)
		}
	}
}

func TestInterface(t *testing.T) {
	var x interface{} = 1
	var y interface{} = []int{1, 2, 3}
	var z interface{} = struct {
		a int
		b byte
	}{1, 'a'}
	tests := []struct {
		in   interface{}
		want string
	}{
		{&x, "(int 1)"},
		{&y, "([]int [1 2 3])"},
		{&z, "(struct { a int; b uint8 } {1 97})"},
	}
	for _, test := range tests {
		actual, _ := Marshal(test.in)
		if string(actual) != test.want {
			t.Errorf("want \"%s\", got \"%s\"", test.want, actual)
		}
	}
}
