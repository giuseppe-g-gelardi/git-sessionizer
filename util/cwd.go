package util

import (
	"fmt"
	"os"
)

func ChangeDir(newDir string) error {
	err := os.Chdir(newDir)
	if err != nil {
		fmt.Printf("Error changing directory: %v\n", err)
		return err
	}

	return nil
}

// likely dont even need to check for curr
// currentDir, err := os.Getwd()
// _, err := os.Getwd()
// if err != nil {
// 	fmt.Printf("Error identifying directory: %v\n", err)
// 	return err
// }

// // newCurrentDir, err := os.Getwd()
// _, err = os.Getwd()
// if err != nil {
//     fmt.Printf("Error: %v\n", err)
//     return err
// }

// cmd := exec.Command("ls")
// cmd.Stdout = os.Stdout
// cmd.Stderr = os.Stderr

// err = cmd.Run()
// if err != nil {
//     return err
// }
