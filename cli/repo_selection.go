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

func setEditorCommand(config *c.Config) string {
	if config.Alias != "" {
		return config.Alias
	}
	return config.Editor
}

var API_URL = "https://api.github.com/user/repos?page={PAGE}&per_page={PER_PAGE}&visibility=all"

func RepoSelection(config *c.Config) {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond) // Build our new spinner])
	s.Start()                                                    // Start the spinner
	s.Suffix = " Fetching all repos..."
	s.Color("cyan")

	var repoUrl string
	editorCmd := setEditorCommand(config)

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

	cmdErr := u.RunCommand([]string{"git", "clone", repoUrl})
	if cmdErr != nil {
		log.Errorf("Error cloning repo: %v", cmdErr)
	}
	cdErr := u.ChangeDir(repo.Name)
	if cdErr != nil {
		log.Errorf("Error changing directory: %v", cdErr)
	}

	if config.Tmux {
		if tmxErr := u.StartTmuxSession(repo.Name, editorCmd); tmxErr != nil {
			log.Errorf("Error starting tmux session: %v", tmxErr)
		}
	}
	if config.Editor == "vscode" {
        cmd := []string{"code", "."}
		if editorErr := u.RunCommand(cmd); editorErr != nil {
			log.Errorf("Error opening editor: %v", editorErr)
		}
	} 

	// the editor also needs to be passed in.
	// if the editor is NOT vim/nvim, just RunCommand([]string{"code", "."}) or whatever
	// if tmxErr := u.StartTmuxSession(repo.Name, editorCmd); tmxErr != nil { // config.Editor
	// 	log.Errorf("Error starting tmux session: %v", tmxErr)
	// }
	// editorErr := u.RunCommand([]string{config.Editor, "."})
	// if editorErr != nil {
	// 	log.Errorf("Error opening editor: %v", editorErr)
	// }
}
