package prompts

import (
	"fmt"

	"github.com/giuseppe-g-gelardi/git-sessionizer/cli/templates"
	"github.com/giuseppe-g-gelardi/git-sessionizer/util"
)

func ConfigureAliasOptions() string {
	var alias string = ""
	description := "An editor alias is a shorthand or custom command used in software development to quickly invoke a specific text editor or integrated development environment (IDE) with predefined settings or options."

	aliasOptions := []templates.DialogOption{
		{
			Name:        "Yes!",
			Value:       true,
			Description: util.WrapText(description, 80),
		},
		{
			Name:        "Nope!",
			Value:       false,
			Description: util.WrapText(description, 80),
		},
	}

	selectedOption := templates.RenderSelect("Confirm Alias Options", aliasOptions, 4)

	if selectedOption == true {
		ans, err := templates.RenderPrompt()
		if err != nil {
			fmt.Printf("Something went wrong... %v\n", err)
		}
		alias = ans
	} else if selectedOption == false {
		alias = ""
	}

	return alias
}
