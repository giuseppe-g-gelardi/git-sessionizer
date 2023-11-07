// package templates

// import (
// 	"strings"

// 	"github.com/manifoldco/promptui"
// )

// type DialogOptions struct {
// 	Name        string `json:"name"`
// 	Value       bool   `json:"value"`
// 	Description string `json:"description"`
// }

// // type PromptTemplate interface {
// // 	SelectTemplates() *promptui.SelectTemplates
// // 	Searcher() func(string, int) bool
// // 	// Prompt(label string, items []DialogOptions, templates *promptui.SelectTemplates searcher func(input string, index int) bool) promptui.Select
// // }

// func SelectTemplates(title string) *promptui.SelectTemplates {
// 	// replace {title} with title
// 	templates := &promptui.SelectTemplates{
// 		Label:    "   {{ .Name | faint }}?",
// 		Active:   "-> {{ .Name | blue }}",
// 		Inactive: "   {{ .Name | cyan }}",
// 		Selected: "   {{ .Name | red | cyan }}",
// 		Details: `
// 		--------- {title} ----------
// {{ "Name:" | faint }}	{{ .Name }}
// {{ "Description:" | faint }}	{{ .Description }}
// 	`,
// 	}
// 	return templates
// }

// func SearcherTemplate(items []DialogOptions) func(string, int) bool {
// 	searcher := func(input string, index int) bool {
// 		option := items[index]
// 		name := strings.Replace(strings.ToLower(option.Name), " ", "", -1)
// 		input = strings.Replace(strings.ToLower(input), " ", "", -1)
// 		return strings.Contains(name, input)
// 	}

// 	return searcher
// }

// func PromptTemplate(label string, items []DialogOptions, templates *promptui.SelectTemplates, searcher func(input string, index int) bool) promptui.Select {
// 	prompt := promptui.Select{
// 		Label:     label,
// 		Items:     items,
// 		Templates: templates,
// 		Size:      4,
// 		Searcher:  searcher,
// 	}
// 	// i, _, err := prompt.Run()
// 	// if err != nil {
// 	// 	fmt.Printf("Prompt failed %v\n", err)
// 	// }
// 	return prompt
// }

package templates

import (
	"fmt"
	"strings"

	u "github.com/giuseppe-g-gelardi/git-sessionizer/util"

	"github.com/manifoldco/promptui"
)

// DialogOption represents an option in a dialog.
type DialogOption struct {
	Name        string
	Value       interface{}
	Description string
}

// RenderPrompt displays a prompt and returns the selected option.
func RenderPrompt(label string, options []DialogOption, maxOptions int) interface{} {
	templates := &promptui.SelectTemplates{
		Label:    "   {{ .Name | faint }}?",
		Active:   "-> {{ .Name | blue }}",
		Inactive: "   {{ .Name | cyan }}",
		Selected: "   {{ .Name | red | cyan }}",
		Details:  `
--------- Repository ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Description:" | faint }}	{{ .Description }}
	`,
	}

	searcher := func(input string, index int) bool {
		option := options[index]
		name := strings.Replace(strings.ToLower(option.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)
		return strings.Contains(name, input)
	}
	prompt := promptui.Select{
		Label:     label,
		Items:     options,
		Templates: templates,
		Size:      maxOptions,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	fmt.Printf("You choose number %v: \n", options[i].Value)

	return options[i].Value
}

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

	return RenderPrompt("Confirm Tmux Options", tmuxOptions, 4).(bool)
}
