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
