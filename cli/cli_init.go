package cli

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/giuseppe-g-gelardi/git-sessionizer/api"
)

var API_URL = "https://api.github.com/user/repos?page={PAGE}&per_page={PER_PAGE}&visibility=all"

func InitCli(token string) {
	// get all user repos
	fmt.Println("Fetching all repos...")
	repos, err := api.FetchAllUserRepos(API_URL, token)
	if err != nil {
		log.Errorf("Error: %v", err)
	}
	fmt.Println("Done!")

	// convert repos to cli.PartialRepo and pass to cli.RepoPrompt to display
	var cliRepos []PartialRepo
	for _, repo := range repos {
		cliRepos = append(cliRepos, PartialRepo(repo))
	}

	RepoPrompt(cliRepos)
}


