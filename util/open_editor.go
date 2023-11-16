package util

import (
	"fmt"
	"os"
	"os/exec"
)

func OpenEditor(editor string) error {
	// editor/alias should be passed in as a string
	cmd := exec.Command(editor, ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error opening %v", editor)
		return err
	}

	return nil
}

