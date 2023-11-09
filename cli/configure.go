package cli

import (
	"fmt"

	p "github.com/giuseppe-g-gelardi/git-sessionizer/cli/prompts"
	"github.com/giuseppe-g-gelardi/git-sessionizer/cli/templates"
	conf "github.com/giuseppe-g-gelardi/git-sessionizer/config"
)

func ConfigureEditor(cm *conf.CfgManager) {

	editor_answer := "nvim"                   // := ConfigureEditorOptions()
	alias_answer := p.ConfigureAliasOptions() // := ConfigureAliasOptions()
	tmux_answer := p.ConfigureTmuxOptions()   // := ConfigureTmuxOptions()

	// ! update the following to return an answer as a bool
	/* conf_answer := */
	ConfirmConfigurationOptions(editor_answer, alias_answer, tmux_answer, cm)
	/*
		if !answer.(bool) {
			ConfigureEditor(cm)
		}
		conf, _ := cm.GetConfig(2)
		InitCli(conf, cm)
	*/
}

func ConfirmConfigurationOptions(editor string, alias string, tmux bool, cm *conf.CfgManager) /* bool */ {
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

	/* return templ... (t/f)*/
	answer := templates.RenderSelect("Confirm Editor Config Options", editorOptions, 4)

	if !answer.(bool) {
		ConfigureEditor(cm)
	}
	conf, _ := cm.GetConfig(2)
	InitCli(conf, cm)
}
