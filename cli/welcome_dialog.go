package cli

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/manifoldco/promptui"
)

func WelcomeDialog() string {
	log.Info("Welcome to Git Sessionizer!")
	Options := []WelcomePromptStruct{
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

	return welcomePrompt(Options)
}

type WelcomePromptStruct struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	Description string `json:"description"`
}

func welcomePrompt(welcomeOptions []WelcomePromptStruct) string {
	templates := &promptui.SelectTemplates{
		Label:    "   {{ .Name | faint }}?",
		Active:   "-> {{ .Name | blue }}",
		Inactive: "   {{ .Name | cyan }}",
		Selected: "   {{ .Name | red | cyan }}",
		Details: `
----------- Welcome ------------
{{ "Description:" | faint }}	{{ .Description }}
`,
	}
	searcher := func(input string, index int) bool {
		option := welcomeOptions[index]
		name := strings.Replace(strings.ToLower(option.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}
	prompt := promptui.Select{
		Label:     "Lets get started",
		Items:     welcomeOptions,
		Templates: templates,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		// return
	}

	// fmt.Printf("You choose number %d: %s\n", i+1, strings.ToLower(welcomeOptions[i].Value))
	return strings.ToLower(welcomeOptions[i].Value)
}

// type PartialRepo struct {
// 	Name        string `json:"name"`
// 	Http_url    string `json:"html_url"`
// 	Ssh_url     string `json:"ssh_url"`
// 	Description string `json:"description"`
// 	Private     bool   `json:"private"`
// }
// func RepoPrompt(repos []PartialRepo) {
// 	templates := &promptui.SelectTemplates{
// 		Label:    "   {{ .Name | faint }}?",
// 		Active:   "-> {{ .Name | blue }} ({{ .Description | red }})",
// 		Inactive: "   {{ .Name | cyan }} ({{ .Description | red }})",
// 		Selected: "   {{ .Name | red | cyan }}",
// 		Details: `
// --------- Repository ----------
// {{ "Name:" | faint }}	{{ .Name }}
// {{ "Description:" | faint }}	{{ .Description }}
// {{ "HTTP URL:" | faint }}	{{ .Http_url }}
// {{ "SSH URL:" | faint }}	{{ .Ssh_url }}
// {{ "Private:" | faint }}	{{ .Private }}
// `,
// 	}

// 	searcher := func(input string, index int) bool {
// 		repo := repos[index]
// 		name := strings.Replace(strings.ToLower(repo.Name), " ", "", -1)
// 		input = strings.Replace(strings.ToLower(input), " ", "", -1)

// 		return strings.Contains(name, input)
// 	}

// 	prompt := promptui.Select{
// 		Label:     "Select a repository",
// 		Items:     repos,
// 		Templates: templates,
// 		Size:      4,
// 		Searcher:  searcher,
// 	}

// 	i, _, err := prompt.Run()

// 	if err != nil {
// 		fmt.Printf("Prompt failed %v\n", err)
// 		return
// 	}

// 	fmt.Printf("You choose number %d: %s\n", i+1, repos[i].Name)
// }
