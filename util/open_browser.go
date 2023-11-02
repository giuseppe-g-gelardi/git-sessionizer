package util

import (
	"fmt"
	"os/exec"
	"runtime"
)

func Openbrowser(url string) error {
	var openCommand string

	switch runtime.GOOS {
	case "darwin":
		openCommand = "open"
	case "linux":
		openCommand = "xdg-open"
	case "windows":
		openCommand = "start"
	default:
		openCommand = "xdg-open"
	}

	cmd := exec.Command(openCommand, url)
	if err := cmd.Start(); err != nil {
		fmt.Printf("Error opening browser: %s\n", err)
		return err
	}

	return nil
}
