package depedencyinjection

import (
	"bytes"
	"testing"
)

// Test this
func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Nishil")

	got := buffer.String()
	want := "Hello Nishil"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
