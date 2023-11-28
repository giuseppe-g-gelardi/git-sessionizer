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

func setEditorCommand(config *c.Config) string {
	if config.Alias != "" {
		return config.Alias
	}
	return config.Editor
}

func setRepoUrl(repo p.PartialRepo) string {

	htmlorssh := p.HtmlOrSsh()

	if htmlorssh == "ssh" {
		return repo.Ssh_url
	}
	return repo.Http_url

}

func RepoSelection(config *c.Config) {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond) // Build our new spinner])
	s.Start()                                                    // Start the spinner
	s.Suffix = " Fetching all repos..."
	s.Color("cyan")

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
	repo, _ := p.RepoPrompt(cliRepos) // this prompt returns the selected repo
	repoUrl := setRepoUrl(repo)       // this returns the repo url
	isBareRepo := p.BareRepoPrompt()
	attach := p.AttachOrStartNewSession()


	
    sessions, _ := u.ListTmuxSessions()
	session_select, _ := p.SessionPrompt(sessions)

	fmt.Printf("session_select: %v\n", session_select)


	fmt.Printf("attach: %v\n", attach)

	c := commandBuilder(repoUrl, isBareRepo)

	// cmdErr := u.RunCommand([]string{"git", "clone", repoUrl})
	cmdErr := u.RunCommand(c)
	if cmdErr != nil {
		log.Errorf("Error cloning repo: %v", cmdErr)
	}

	startSession(repo, config, editorCmd)
}

func commandBuilder(repoUrl string, isBareRepo bool) []string {
	var cmd []string
	// so janky
	if isBareRepo {
		cmd = []string{"git", "clone", "--bare", repoUrl}
	} else {
		cmd = []string{"git", "clone", repoUrl}
	}

	return cmd
}

func startSession(repo p.PartialRepo, config *c.Config, editorCmd string) {
	cdErr := u.ChangeDir(repo.Name)
	if cdErr != nil {
		log.Errorf("Error changing directory: %v", cdErr)
	}

	// clean this upppppppppp
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
}
