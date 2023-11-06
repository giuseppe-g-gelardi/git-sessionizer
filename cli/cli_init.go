package cli

import (
	u "github.com/giuseppe-g-gelardi/git-sessionizer/util"
	conf "github.com/giuseppe-g-gelardi/git-sessionizer/config"
)

func InitCli(token string, cm *conf.ConfigManager) {
	// clear the console first? // conf should be the configManager?
	welcome := WelcomeDialog()

	switch welcome {
	case "open":
		RepoSelection(token /* conf */)
	case "update":
		ConfigureEditor(cm)
	case "exit":
		u.Exit()
	}
}
