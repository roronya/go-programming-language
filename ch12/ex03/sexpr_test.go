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
	for _, t := range tests {
		actual, _ := Marshal(t.in)
		if string(actual) != t.want {
			log.Fatalf("want \"%s\", got \"%s\"", t.want, actual)
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
	}
	for _, t := range tests {
		actual, _ := Marshal(t.in)
		if string(actual) != t.want {
			log.Fatalf("want \"%s\", got \"%s\"", t.want, actual)
		}
	}
}
