package prompts

import (
	"github.com/giuseppe-g-gelardi/git-sessionizer/cli/templates"
	"github.com/giuseppe-g-gelardi/git-sessionizer/util"
)

func BareRepoPrompt() bool {
	var cloneOption bool = false

	cloneOptions := []templates.DialogOption{
		{
			Name:        "No (recommended)",
			Value:       false,
			Description: util.WrapText("regular, please!", 80),
		},
		{
			Name:        "Yes",
			Value:       true,
			Description: util.WrapText("Clones the repository as a bare repository -- great for worktrees and .dotfiles", 80),
		},
	}

	selectedOption := templates.RenderSelect("Clone as 'bare' or regular repository?", cloneOptions, 4)

	if selectedOption == false {
		cloneOption = false
	} else if selectedOption == true {
		cloneOption = true
	}

	return cloneOption
}
