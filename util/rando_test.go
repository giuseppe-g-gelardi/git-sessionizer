package util

import (
	"testing"
)

func TestRando(t *testing.T) {
	// Test with no arguments
	result := Rando()
	if result < 0 || result > 3 {
		t.Errorf("Rando() = %d; want a value between 0 and 3", result)
	}

	// Test with one argument
	result = Rando(10)
	if result < 0 || result > 9 {
		t.Errorf("Rando(10) = %d; want a value between 0 and 9", result)
	}

	// Test with multiple arguments (should only use the first argument)
	result = Rando(10, 20, 30)
	if result < 0 || result > 9 {
		t.Errorf("Rando(10, 20, 30) = %d; want a value between 0 and 9", result)
	}
}
