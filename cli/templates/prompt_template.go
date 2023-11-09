package templates

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func RenderPrompt() (string, error) {

	prompt := promptui.Prompt{
		Label: "Enter your alias",
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	return result, nil
}
