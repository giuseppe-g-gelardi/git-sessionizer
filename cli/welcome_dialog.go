package cli

import "github.com/giuseppe-g-gelardi/git-sessionizer/cli/templates"

type WelcomePromptStruct struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	Description string `json:"description"`
}

func WelcomeDialog() string {
	welcomeOptions := []templates.DialogOption{
		{
			Name:        "Open a repo",
			Value:       "open",
			Description: "I'd like to clone and open a repo",
		},
		{
			Name:        "Update my config",
			Value:       "update",
			Description: "Join a session",
		},
		{
			Name:        "Exit",
			Value:       "exit",
			Description: "Exit the application",
		},
	}

	return templates.RenderPrompt("Welcome to Git Sessionizer!", welcomeOptions, 4).(string)
}
