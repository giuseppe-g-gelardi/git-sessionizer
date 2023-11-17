package util

import (
	"fmt"
	"os"
	"os/exec"
)

func RunCommand(command []string) error {
	// func RunCommand(command ...string) error {
	fmt.Printf(command[0], command[1:])

	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// err := cmd.Run()
	//    // output, err := cmd.CombinedOutput()
	// if err != nil {
	//        fmt.Printf("Error running command: %v", command)
	// 	return err
	// }

	// fmt.Printf("Output: %v", output)
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

    fmt.Printf("sessionName: %v", sn)
    fmt.Println("tmux", "new", "-s", string(sn))

	// tmuxCmd := exec.Command("tmux", "new", "-s", "devvvvv")
	tmuxCmd := exec.Command("tmux", "new", "-s", string(sn))
	tmuxCmd.Stdout = os.Stdout
	tmuxCmd.Stderr = os.Stderr
	tmuxCmd.Stdin = os.Stdin


	// nvimCmd := exec.Command("nvim .")
	// nvimCmd.Stdout = os.Stdout
	// nvimCmd.Stderr = os.Stderr
	// nvimCmd.Stdin = os.Stdin

	tmuxErr := tmuxCmd.Start()
	if tmuxErr != nil {
		fmt.Printf("Error running tmux command: %v", tmuxErr)
		return tmuxErr
	}

	// nvimErr := nvimCmd.Start()
	// if nvimErr != nil {
	// 	fmt.Printf("Error running nvim command: %v", nvimErr)
	// 	return nvimErr
	// }

	tmuxErr = tmuxCmd.Wait()
	if tmuxErr != nil {
		fmt.Printf("Error running tmux command: %v", tmuxErr)
		return tmuxErr
	}
	// nvimErr = nvimCmd.Wait()
	// if nvimErr != nil {
	// 	fmt.Printf("Error running nvim command: %v", nvimErr)
	// 	return nvimErr
	// }

	return nil

}
