package cli

import (
	"fmt"
	"time"

	"github.com/giuseppe-g-gelardi/git-sessionizer/api"
	p "github.com/giuseppe-g-gelardi/git-sessionizer/cli/prompts"

	"github.com/briandowns/spinner"
	"github.com/charmbracelet/log"
)

var API_URL = "https://api.github.com/user/repos?page={PAGE}&per_page={PER_PAGE}&visibility=all"

func RepoSelection(token string) {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond) // Build our new spinner])
	s.Start()                                                    // Start the spinner
	s.Suffix = " Fetching all repos..."
	s.Color("cyan")
	repos, err := api.RepoList(API_URL, token)
	if err != nil {
		log.Errorf("Error: %v", err)
	}
	s.Stop() // Stop the spinner

	// convert repos to cli.PartialRepo and pass to cli.RepoPrompt to display
	var cliRepos []p.PartialRepo
	for _, repo := range repos {
		cliRepos = append(cliRepos, p.PartialRepo(repo))
	}
	repo, _ := p.RepoPrompt(cliRepos)

	htmlorssh := p.HtmlOrSsh()

	if htmlorssh == "https" {
		fmt.Printf("You chose HTTPS\n")
		fmt.Printf("CLONEURL %v\n", repo.Http_url)
	} else if htmlorssh == "ssh" {
		fmt.Printf("You chose SSH\n")
		fmt.Printf("CLONEURL %v\n", repo.Ssh_url)
	}

	/*
	   clone repo (via https or ssh)
	   cd into (repo.Name)
       
       compile and run list of commands {
           tmux, editor, alias, etc.
       }
	*/
}
