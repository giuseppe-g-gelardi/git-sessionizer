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
	cmd.Stdin = os.Stdin

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

func RunTmuxAndNvim(sessionName string) error {
	sn := StrFormat(sessionName)

	tmuxCmd := exec.Command("tmux", "new", "-s", string(sn))
	tmuxCmd.Stdout = os.Stdout
	tmuxCmd.Stderr = os.Stderr
	tmuxCmd.Stdin = os.Stdin

	tmuxErr := tmuxCmd.Start()
	if tmuxErr != nil {
		fmt.Printf("Error running tmux command: %v", tmuxErr)
		return tmuxErr
	}

	tmuxErr = tmuxCmd.Wait()
	if tmuxErr != nil {
		fmt.Printf("Error running tmux command: %v", tmuxErr)
		return tmuxErr
	}

	nvimErr := RunCommand([]string{"tmux", "send-keys", "-t", string(sn), "nvim", ".", "C-m"})
	if nvimErr != nil {
		fmt.Printf("Error running nvim command: %v", nvimErr)
		return nvimErr
	}

	return nil

}

func StartTmuxSession(sessionName string) error {
	// session := strings.ReplaceAll(sessionName, ".", "_")
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
    // THIS IS CRUTIAL AND ALSO A HACK
	time.Sleep(1 * time.Second)

	// Send keys to open nvim within the tmux session
	nvimCmd := exec.Command("tmux", "send-keys", "-t", session, "nvim .", "C-m")
	nvimCmd.Stdout = os.Stdout
	nvimCmd.Stderr = os.Stderr

	if err := nvimCmd.Run(); err != nil {
		return fmt.Errorf("error sending keys to tmux: %v", err)
	}

	// Wait for tmux to finish
	if err := tmuxCmd.Wait(); err != nil {
		return fmt.Errorf("error waiting for tmux: %v", err)
	}

	return nil
}

