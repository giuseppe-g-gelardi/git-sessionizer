package cli

import (
	"fmt"

	"github.com/giuseppe-g-gelardi/git-sessionizer/cli/templates"
	conf "github.com/giuseppe-g-gelardi/git-sessionizer/config"
)

func ConfigureEditor(cm *conf.ConfigManager) {

	editor_answer := "nvim"                 // := ConfigureEditorOptions()
	alias_answer := ConfigureAliasOptions() // := ConfigureAliasOptions()
	tmux_answer := ConfigureTmuxOptions()   // := ConfigureTmuxOptions()

	confirmEditorOptions(editor_answer, alias_answer, tmux_answer, cm)
}

func confirmEditorOptions(editor string, alias string, tmux bool, cm *conf.ConfigManager) {
	fmt.Println("Your config options are:")
	fmt.Printf("Editor: %s\n", editor)
	if tmux {
		fmt.Printf("Tmux: %t\n", tmux)
	}
	if alias != "" {
		fmt.Printf("Alias: %v\n", alias)
	}

	editorOptions := []templates.DialogOption{
		{
			Name:        "Yes!",
			Value:       true,
			Description: "I'm happy with these options",
		},
		{
			Name:        "Nope!",
			Value:       false,
			Description: "I'd like to update my config",
		},
	}
	answer := templates.RenderSelect("Confirm Editor Config Options", editorOptions, 4)
	if !answer.(bool) {
		ConfigureEditor(cm)
	}
	conf, _ := cm.GetConfig(2)
	InitCli(conf, cm)
}
