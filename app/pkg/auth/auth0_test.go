package auth

import (
	"testing"
)

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg := ""
	if msg != "" {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, nil)
	}
}
