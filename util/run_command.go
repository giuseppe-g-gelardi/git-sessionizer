package util

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func RunCommand(command []string) error {

	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		fmt.Printf("Error running command: %v", command)
		return err
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Printf("Error running command: %v", command)
		return err
	}

	return nil
}

func StartTmuxSession(sessionName string, editorCmd string) error {
	editorCmd = editorCmd + " ."
	session := StrFormat(sessionName)

	// Start the tmux session
	tmuxCmd := exec.Command("tmux", "new", "-s", string(session))
	tmuxCmd.Stdout = os.Stdout
	tmuxCmd.Stderr = os.Stderr
	tmuxCmd.Stdin = os.Stdin

	if err := tmuxCmd.Start(); err != nil {
		return fmt.Errorf("error starting tmux: %v", err)
	}

	// Wait for a moment to allow the tmux session to initialize
	time.Sleep(1 * time.Second) // THIS IS CRUTIAL AND ALSO A HACK LOL

	// Send keys to open nvim within the tmux session
	nvimErr := RunCommand([]string{"tmux", "send-keys", "-t", string(session), editorCmd, "C-m"})
	if nvimErr != nil {
		return nvimErr
	}

	// Wait for tmux to finish
	if err := tmuxCmd.Wait(); err != nil {
		return fmt.Errorf("error waiting for tmux: %v", err)
	}

	return nil
}

// // maybe...
// func run(command []string) (*exec.Cmd, error) {
// 	cmd := exec.Command(command[0], command[1:]...)
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr
// 	cmd.Stdin = os.Stdin
//
// 	if err := cmd.Run(); err != nil {
// 		return nil, fmt.Errorf("error running command %v: %v", command, err)
// 	}
//
// 	return cmd, nil
// }
