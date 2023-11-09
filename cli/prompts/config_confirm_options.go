package prompts

import (
	"fmt"

	"github.com/giuseppe-g-gelardi/git-sessionizer/cli/templates"
	conf "github.com/giuseppe-g-gelardi/git-sessionizer/config"
)

func ConfirmConfigurationOptions(editor string, alias string, tmux bool, cm *conf.CfgManager) bool {
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

	return templates.RenderSelect("Confirm Editor Config Options", editorOptions, 4).(bool)
}
