package prompts

import (
	"github.com/giuseppe-g-gelardi/git-sessionizer/cli/templates"
	u "github.com/giuseppe-g-gelardi/git-sessionizer/util"
)

func ConfigureEditorOptions() string {
	var vimDescriptionOptions = []string{
		`Vim is a highly efficient and configurable text editor known for its modal editing system and extensive keyboard shortcuts.`,
		`Vim is like the Rubik's Cube of text editors – perplexing at first, but once you figure it out, you'll feel like a wizard.`,
		`If you're into memorizing cryptic key combinations and enjoy feeling like you're trapped in the '80s, Vim might be your text editor of choice.`,
	}

	var neovimDescriptionOptions = []string{
		`Neovim is a modernized and extended version of Vim, designed to be more extensible and maintainable while retaining Vim's core features and compatibility.`,
		`Neovim is like Vim's hipster cousin – it does all the same things but claims it was doing them before they were cool.`,
		`Does your linked in say "open to work" and you spend all day watching thePrimeagen?`,
	}

	var vscodeDescriptionOptions = []string{
		`Visual Studio Code (VSCode) is a code editor that caters to those who enjoy a more visually cluttered and resource-intensive development environment.`,
		`Visual Studio Code (VSCode) is a code editor for people who prefer a user-friendly, feature-rich, and less keyboard-driven development experience.`,
		`Visual Studio Code (VSCode) is like Vim's younger sibling – it's not as powerful, but it's easier to use and more fun to play with.`,
	}

	vimDesc := vimDescriptionOptions[u.Rando(3)]
	nvimDesc := neovimDescriptionOptions[u.Rando(3)]
	vscodeDesc := vscodeDescriptionOptions[u.Rando(3)]

    editorOptions := []templates.DialogOption{
        {
            Name:        "Vim",
            Value:       "vim",
            Description: vimDesc,
        },
        {
            Name:        "Neovim",
            Value:       "nvim",
            Description: nvimDesc,

        },
        {
            Name:        "VSCode",
            Value:       "vscode",
            Description: vscodeDesc,
        },
    }

	return templates.RenderSelect("Configure Editor", editorOptions, 4).(string)
}
