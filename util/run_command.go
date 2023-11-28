package util

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

    // p "github.com/giuseppe-g-gelardi/git-sessionizer/cli/prompts"
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
    

    // attach := p.AttachOrStartNewSession()
    // sessions, _ := listTmuxSessions()
    // session_select, _ := p.SessionPrompt(sessions)




	/*
	*/

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

func ListTmuxSessions() ([]string, error) {

	active, err := isTmuxActive()
	if err != nil {
		return nil, err
	}

	if !active {
		return nil, fmt.Errorf("tmux is not active")
	}

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



func isTmuxActive() (bool, error) {
	cmd := exec.Command("tmux", "info")
	cmd.Stderr = cmd.Stdout

	err := cmd.Run()

	return err == nil, nil
}
