package main

import (
	"fmt"

	"github.com/manifoldco/promptui"

	"github.com/giuseppe-g-gelardi/git-sessionizer/auth"
)

func main() {
	auth.Authenticate()


	prompt_ui()
}

func prompt_ui() {
	// Define a list of options
	items := []string{"Apple", "Banana", "Orange", "Grapes", "Strawberry"}

	// Create a Select prompt
	prompt := promptui.Select{
		Label: "Select a fruit",
		Items: items,
	}

	// Show the prompt to the user
	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You selected: %s\n", result)

}
