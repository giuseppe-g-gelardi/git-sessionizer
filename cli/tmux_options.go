package cli

import (
	"fmt"
	"strings"

	u "github.com/giuseppe-g-gelardi/git-sessionizer/util"

	"github.com/manifoldco/promptui"
)

func ConfigureTmuxOptions() bool {
	var tmuxDescriptionOptions = []string{
		"Tmux is a powerful terminal multiplexer that enables efficient session management, allowing you to create and switch between multiple terminal panes and windows.",
		"Tmux is like the command-line's Swiss Army knife, where you can slice, dice, and juggle your terminal tasks with finesse.",
		"Tmux is for devs that enjoy fumbling through multiple terminal windows and never knowing which one has your code.",
	}

	desc := tmuxDescriptionOptions[u.Rando(3)]

	tmuxOptions := []DialogOption{
		{
			Name:        "Yes!",
			Value:       true,
			Description: desc,
		},
		{
			Name:        "Nope!",
			Value:       false,
			Description: desc,
		},
	}
	templates := &promptui.SelectTemplates{
		Label:    "   {{ .Name | faint }}?",
		Active:   "-> {{ .Name | blue }}",
		Inactive: "   {{ .Name | cyan }}",
		Selected: "   {{ .Name | red | cyan }}",
		Details: `
--------- Repository ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Description:" | faint }}	{{ .Description }}
	`,
	}

	searcher := func(input string, index int) bool {
		option := tmuxOptions[index]
		name := strings.Replace(strings.ToLower(option.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)
		return strings.Contains(name, input)
	}
	prompt := promptui.Select{
		Label:     "Confirm Tmux Options",
		Items:     tmuxOptions,
		Templates: templates,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	// // cfgCurr, err := cm.GetConfig(1)
	// // fmt.Printf("Current Config: %v", cfgCurr)

	// cfg, err := cm.GetConfig(1)
	// if err != nil {
	// 	fmt.Printf("Error getting config: %v", err)
	// }

	// // this should bring you back to the initCli function
	// fmt.Printf("You choose number %d: %s\n", i+1, editorOptions[i].Name)
	// InitCli(cfg, cm)

	fmt.Printf("You choose number %v: \n", tmuxOptions[i].Value)

	return tmuxOptions[i].Value
}
