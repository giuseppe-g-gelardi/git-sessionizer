package cli

import (
	conf "github.com/giuseppe-g-gelardi/git-sessionizer/config"
	u "github.com/giuseppe-g-gelardi/git-sessionizer/util"
)

func InitCli(config *conf.UserConfig, cm *conf.ConfigManager) {
	// clear the console first?
	welcome := WelcomeDialog()

	switch welcome {
	case "open":
		RepoSelection(config.AccessToken)
	case "update":
		ConfigureEditor(cm)
	case "exit":
		u.Exit()
	}
}
