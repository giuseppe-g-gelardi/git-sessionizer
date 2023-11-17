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
