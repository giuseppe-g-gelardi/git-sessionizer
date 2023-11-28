package util

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
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

func AttachTmuxSession(sessionName, editorCmd string) error {
	editorCmd = editorCmd + " ."
	session := StrFormat(sessionName)

	tmuxCmd := exec.Command("tmux", "new-window", "-t", session)
	tmuxCmd.Stdout = os.Stdout
	tmuxCmd.Stderr = os.Stderr
	tmuxCmd.Stdin = os.Stdin

	// create the new tmux window
	if err := tmuxCmd.Start(); err != nil {
		return fmt.Errorf("error starting tmux: %v", err)
	}

	// get the list of windows in the selected tmux session
	windows := exec.Command("tmux", "list-windows", "-t", session)
	windows.Stderr = os.Stderr
	out, err := windows.Output()

	var window_names []string

	if err != nil {
		return err
	}

	// iterate through the list of windows and get the window names
	windows_list := strings.Split(string(out), "\n")
	for _, window := range windows_list {
		window_names = append(window_names, strings.Split(window, ":")[0])
	}

	// send keys to the last window in the list to open editor (new window seems to always be the last)
	winErr := RunCommand([]string{"tmux", "send-keys", "-t", session + ":" + window_names[len(window_names)-1], editorCmd, "C-m"})
	if winErr != nil {
		return winErr
	}

	if err := tmuxCmd.Wait(); err != nil {
		return fmt.Errorf("error waiting for tmux: %v", err)
	}
	return nil
}

func StartTmuxSession(sessionName, editorCmd string) error {
	editorCmd = editorCmd + " ."
	session := StrFormat(sessionName)

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

func ListTmuxSessions() ([]string, error) {

	// active, err := IsTmuxActive()
	// if err != nil {
	// 	return nil, err
	// }
	//
	// if !active {
	// 	return nil, fmt.Errorf("tmux is not active")
	// }

	cmd := exec.Command("tmux", "list-sessions")
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()

	var session_names []string

	if err != nil {
		return nil, err
	}

	sessions := strings.Split(string(out), "\n")
	for _, session := range sessions {
		session_names = append(session_names, strings.Split(session, ":")[0])
	}

	return session_names, nil
}

func IsTmuxActive() (bool, error) {
	cmd := exec.Command("tmux", "info")
	cmd.Stderr = cmd.Stdout

	err := cmd.Run()

	return err == nil, nil
}
