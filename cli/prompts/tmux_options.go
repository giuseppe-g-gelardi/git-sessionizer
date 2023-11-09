package prompts

import (
	"github.com/giuseppe-g-gelardi/git-sessionizer/cli/templates"
	u "github.com/giuseppe-g-gelardi/git-sessionizer/util"
)

func ConfigureTmuxOptions() bool {
	var tmuxDescriptionOptions = []string{
		"Tmux is a powerful terminal multiplexer that enables efficient session management, allowing you to create and switch between multiple terminal panes and windows.",
		"Tmux is like the command-line's Swiss Army knife, where you can slice, dice, and juggle your terminal tasks with finesse.",
		"Tmux is for devs that enjoy fumbling through multiple terminal windows and never knowing which one has your code.",
	}

	desc := tmuxDescriptionOptions[u.Rando(3)]

	tmuxOptions := []templates.DialogOption{
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

	return templates.RenderSelect("Confirm Tmux Options", tmuxOptions, 4).(bool)
}
