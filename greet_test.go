package learn

/*
import "testing"
import "fmt"
*/

import (
	"testing"
)

func TestGreet(t *testing.T) {
	// english
	got := Greet("en", "World")
	exp := "Hi, World!"
	if got != exp {
		t.Errorf("Expected %q, got %q", exp, got)
	}

	// svenska
	got = Greet("sv", "World")
	exp = "Hej, World!"
	if got != exp {
		t.Errorf("Expected %q, got %q", exp, got)
	}
}
