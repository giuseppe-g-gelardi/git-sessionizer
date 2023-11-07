package templates

import (
	"fmt"
	"math"
	"strings"

	"github.com/manifoldco/promptui"
)

// DialogOption represents an option in a dialog.
type DialogOption struct {
	Name        string
	Value       interface{}
	Description string
}

// RenderTitle generates a centered title line based on the given title and line length.
func RenderTitle(title string, lineLength int) string {
	padding := int(math.Max(0, float64(lineLength-len(title))) / 2)
	line := strings.Repeat("-", padding) + title + strings.Repeat("-", lineLength-padding-len(title))
	return line
}

// RenderPrompt displays a prompt and returns the selected option.
func RenderPrompt(label string, options []DialogOption, maxOptions int) interface{} {
	templates := &promptui.SelectTemplates{
		Label:    "   {{ .Name | faint }}?",
		Active:   "-> {{ .Name | blue }}",
		Inactive: "   {{ .Name | cyan }}",
		Selected: "   {{ .Name | red | cyan }}",
		Details: fmt.Sprintf(`
%s
{{ "Name:" | faint }}	{{ .Name }}
{{ "Description:" | faint }}	{{ .Description }}
`, RenderTitle(label, 80)),
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
