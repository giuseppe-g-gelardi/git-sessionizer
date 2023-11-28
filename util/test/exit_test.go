package util

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestExit(t *testing.T) {
	// Redirect stdout to a buffer
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run the Exit function
	Exit()

	// Restore stdout
	w.Close()
	os.Stdout = old

	// Read the output from the buffer
	var buf bytes.Buffer
	io.Copy(&buf, r)

	// Check that the output contains "Bye!"
	if !strings.Contains(buf.String(), "Bye!") {
		t.Errorf("Exit() did not print 'Bye!': %s", buf.String())
	}
}
