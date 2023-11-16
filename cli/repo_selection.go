package cli

import (
	"fmt"
	"time"

	"github.com/giuseppe-g-gelardi/git-sessionizer/api"
	p "github.com/giuseppe-g-gelardi/git-sessionizer/cli/prompts"
	c "github.com/giuseppe-g-gelardi/git-sessionizer/config"
	u "github.com/giuseppe-g-gelardi/git-sessionizer/util"

	"github.com/briandowns/spinner"
	"github.com/charmbracelet/log"
)

var API_URL = "https://api.github.com/user/repos?page={PAGE}&per_page={PER_PAGE}&visibility=all"

func RepoSelection(config *c.Config) {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond) // Build our new spinner])
	s.Start()                                                    // Start the spinner
	s.Suffix = " Fetching all repos..."
	s.Color("cyan")

	var repoUrl string

	repos, err := api.RepoList(API_URL, config.AccessToken)
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
		repoUrl = repo.Http_url
	} else if htmlorssh == "ssh" {
		fmt.Printf("You chose SSH\n")
		fmt.Printf("CLONEURL %v\n", repo.Ssh_url)
		repoUrl = repo.Ssh_url
	}

	/*
			   clone repo (via https or ssh)
			   cd into (repo.Name)

		       compile and run list of commands {
		           tmux, editor, alias, etc.
		       }
	*/

	// // open editor logic:
	/*
	   var editorCmd string
	   if config.Alias != "" {
	       editorCmd = config.Alias
	   }
	   editorCmd = config.Editor
	*/

	/*
			   list of commands: [
		       clone: ["git", "clone", repoUrl],
		       cdDir: ["cd", repo.Name],
		       code: [config.Editor, "."],
			   ]
	*/

	cmdErr := u.RunCommand([]string{"git", "clone", repoUrl})
	// cmdErr := u.RunCommand("git", "clone", string(repoUrl))
	if cmdErr != nil {
		log.Errorf("Error cloning repo: %v", cmdErr)
	}
	cdErr := u.ChangeDir(repo.Name)
	if cdErr != nil {
		log.Errorf("Error changing directory: %v", cdErr)
	}
	// fmtStr := u.StrFormat(repo.Name)
	// fmt.Printf("fmtStr: %v\n", fmtStr)

	// tmuxErr := u.RunCommand([]string{"tmux", "new", "-S", repo.Name})
	tmuxErr := u.RunCommand([]string{"tmux"})
	if tmuxErr != nil {
		log.Errorf("Error creating tmux session: %v", tmuxErr)
	}

	/*
		perhaps the command to open the editor needs to be run in the same shell as the tmux sessions
		take notes from the ts impl
	*/

	// editorErr := u.RunCommand([]string{config.Editor, "."})
	// if editorErr != nil {
	// 	log.Errorf("Error opening editor: %v", editorErr)
	// }

	/*
	  THIS IS SO MUCH EASIER IN GO
	*/

}
