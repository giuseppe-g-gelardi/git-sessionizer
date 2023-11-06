package cli

import (
	// "fmt"
    "time"

	"github.com/giuseppe-g-gelardi/git-sessionizer/api"

    "github.com/charmbracelet/log"
    "github.com/briandowns/spinner"
)

var API_URL = "https://api.github.com/user/repos?page={PAGE}&per_page={PER_PAGE}&visibility=all"

func InitCli(token string, url string) {
    s := spinner.New(spinner.CharSets[14], 100*time.Millisecond) // Build our new spinner])
    s.Start() // Start the spinner
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


