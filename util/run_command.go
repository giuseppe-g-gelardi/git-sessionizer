package util

import (
	"fmt"
	"os"
	"os/exec"
)

func RunCommand(command []string) error {
    fmt.Printf(command[0], command[1:])


	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
        fmt.Printf("Error running command: %v", command)
		return err
	}

	return nil
}
