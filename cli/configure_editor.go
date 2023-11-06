package cli

import (
	"fmt"
	"strings"

	conf "github.com/giuseppe-g-gelardi/git-sessionizer/config"

	"github.com/manifoldco/promptui"
)

type DialogOption struct {
	Name        string `json:"name"`
	Value       bool   `json:"value"`
	Description string `json:"description"`
}

func ConfigureEditor(cm *conf.ConfigManager) {
	// something like this:
	editor_answer := "nvim" // := ConfigureEditorOptions()
	alias_answer := false   //:= ConfigureAliasOptions()
	tmux_answer := true     // := ConfigureTmuxOptions()

	fmt.Println("this is the confirmation dialog?")
	// should just use the config or bring in the config manager

	confirmEditorOptions(editor_answer, alias_answer, tmux_answer, cm)
}

func confirmEditorOptions(editor string, alias bool, tmux bool, cm *conf.ConfigManager) {
	fmt.Printf("Confirm Editor Options %v %v %v", editor, alias, tmux)
	editorOptions := []DialogOption{
		{
			Name:        "LGTM!",
			Value:       true,
			Description: "I'm happy with these options",
		},
		{
			Name:        "Update my config",
			Value:       false,
			Description: "I'd like to update my config",
		},
	}
	// editorPrompt(Options)
	templates := &promptui.SelectTemplates{
		Label:    "   {{ .Name | faint }}?",
		Active:   "-> {{ .Name | blue }} ({{ .Description | red }})",
		Inactive: "   {{ .Name | cyan }} ({{ .Description | red }})",
		Selected: "   {{ .Name | red | cyan }}",
		Details: `
	--------- Repository ----------
	{{ "Name:" | faint }}	{{ .Name }}
	{{ "Description:" | faint }}	{{ .Description }}
	`,
	}

	searcher := func(input string, index int) bool {
		repo := editorOptions[index]
		name := strings.Replace(strings.ToLower(repo.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Select a repository",
		Items:     editorOptions,
		Templates: templates,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if !editorOptions[i].Value {
		fmt.Println("this *SHOULD* bring you back to the config editor")
	}

	cfgCurr, err := cm.GetConfig(1)
	if err != nil {
		fmt.Printf("Error getting config: %v", err)
	}
	fmt.Printf("Current Config: %v", cfgCurr)

	// this should bring you back to the initCli function
	fmt.Printf("You choose number %d: %s\n", i+1, editorOptions[i].Name)
	// InitCli("", &conf.UserConfig{})
}
