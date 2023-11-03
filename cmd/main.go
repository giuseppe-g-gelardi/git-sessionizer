package main

import (
	"fmt"

	"github.com/manifoldco/promptui"

	"github.com/giuseppe-g-gelardi/git-sessionizer/auth"
	c "github.com/giuseppe-g-gelardi/git-sessionizer/config"
)

func main() {

	cfg, err := c.NewConfigManager().GetConfig(3)
	if err != nil {
		fmt.Println(err)
		return
	}

	if cfg.AccessToken == "" {
		fmt.Println("No access token found")

		token, err := auth.Authenticate()
		if err != nil {
			fmt.Println("Error authenticating")
			return
		}
		// cfg.AccessToken = `"token": "` + token
		cfg.AccessToken = token
		if _, err := c.NewConfigManager().WriteConfig(cfg); err != nil {
			fmt.Println("Error writing config")
			return
		}
	} else {
		fmt.Println("Access token found")
	}

	uCfg := cfg

	if _, err := c.NewConfigManager().WriteConfig(uCfg); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Config: %+v\n", uCfg)
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
