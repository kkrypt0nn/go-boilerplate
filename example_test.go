package boilerplate

import (
	"testing"
)

func TestExample42(t *testing.T) {
	expected := 42
	got := Example42()
	if expected != got {
		t.Errorf("Got invalid return value (expected: %d, got %d)", expected, got)
		return
	}

	t.Log("Successfully got 42")
}
