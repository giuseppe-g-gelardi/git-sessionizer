package cli

import (
	"github.com/giuseppe-g-gelardi/git-sessionizer/cli/templates"
)

func ConfigureAliasOptions() string {
	var alias string = ""

	aliasOptions := []templates.DialogOption{
		{
			Name:        "Yes!",
			Value:       true,
			Description: "an editor alias is a shorthand or custom command used in software development to quickly invoke a specific text editor or integrated development environment (IDE) with predefined settings or options.",
		},
		{
			Name:        "Nope!",
			Value:       false,
			Description: "No, I open with the standard command",
		},
	}

	selectedOption := templates.RenderPrompt("Confirm Alias Options", aliasOptions, 4)

	if selectedOption == true {
		/*
			-- setup input prompt for user to enter their alias
			-- alias = `input prompt` -- "c ." or "v ." <-- something like that
		*/
	} else if selectedOption == false {
		alias = ""
	}

	// return templates.RenderPrompt("Confirm Alias Options", alias, 4).(string)
	return alias
}
