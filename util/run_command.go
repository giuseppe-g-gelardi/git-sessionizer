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
