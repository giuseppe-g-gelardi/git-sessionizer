package cli

import (
	"time"

	"github.com/giuseppe-g-gelardi/git-sessionizer/api"
	u "github.com/giuseppe-g-gelardi/git-sessionizer/util"

	"github.com/briandowns/spinner"
	"github.com/charmbracelet/log"
)

var API_URL = "https://api.github.com/user/repos?page={PAGE}&per_page={PER_PAGE}&visibility=all"

func InitCli(token string, url string) {

	welcome := WelcomeDialog()

	switch welcome {
	case "open":
		RepoDialog(token, url)
	case "update":
		log.Infof("!!!!You chose: %s", welcome)
	case "exit":
		u.Exit()
	}

}

func RepoDialog(token string, url string) {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond) // Build our new spinner])
	s.Start()                                                    // Start the spinner
	s.Suffix = " Fetching all repos..."
	s.Color("cyan")
	repos, err := api.FetchAllUserRepos(url, token)
	if err != nil {
		log.Errorf("Error: %v", err)
	}
	s.Stop() // Stop the spinner

	// convert repos to cli.PartialRepo and pass to cli.RepoPrompt to display
	var cliRepos []PartialRepo
	for _, repo := range repos {
		cliRepos = append(cliRepos, PartialRepo(repo))
	}
	RepoPrompt(cliRepos)
}
