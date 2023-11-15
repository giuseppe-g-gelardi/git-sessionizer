package cli

import (
	p "github.com/giuseppe-g-gelardi/git-sessionizer/cli/prompts"
	conf "github.com/giuseppe-g-gelardi/git-sessionizer/config"
	u "github.com/giuseppe-g-gelardi/git-sessionizer/util"
)

func InitCli(config *conf.Config, cm *conf.CfgManager) {
	// clear the consone when this function gets called/reacalled
	u.Clear()
	// display the welcome dialog options
	welcome := p.WelcomeDialog()


	// switch on the welcome dialog options:
	switch welcome {
	case "open": // open a repo,
		RepoSelection(config.AccessToken)
	case "update": // update the editor config,
		ConfigureEditor(cm)
	case "exit": // or exit the program
		u.Exit()
	}
}
