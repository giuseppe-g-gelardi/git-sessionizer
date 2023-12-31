package prompts

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

type PartialRepo struct {
	Name        string `json:"name"`
	Http_url    string `json:"html_url"`
	Ssh_url     string `json:"ssh_url"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
}

func RepoPrompt(repos []PartialRepo) (PartialRepo, error) {
	templates := &promptui.SelectTemplates{
		Label:    "   {{ .Name | faint }}?",
		Active:   "-> {{ .Name | blue }}",
		Inactive: "   {{ .Name | cyan }}",
		Selected: "   {{ .Name | red | cyan }}",
		Details: `
--------- Repository ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Description:" | faint }}	{{ .Description }}
{{ "HTTP URL:" | faint }}	{{ .Http_url }}
{{ "SSH URL:" | faint }}	{{ .Ssh_url }}
{{ "Private:" | faint }}	{{ .Private }}
`,
	}

	searcher := func(input string, index int) bool {
		repo := repos[index]
		name := strings.Replace(strings.ToLower(repo.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Select a repository",
		Items:     repos,
		Templates: templates,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		// return _, errors.New("Prompt failed")
	}

	// fmt.Printf("You chose repo %d: %s\n", i+1, repos[i].Name)
	//    fmt.Printf("http url: %s\n", repos[i].Http_url)
	//    fmt.Printf("ssh url: %s\n", repos[i].Ssh_url)

	// should probably return the whole *partial* repo object just incase the
	// user wants to clone ssh OR html.
	// - will probably just start with http implementation

	// return // repos[i].Name

	return repos[i], nil
}
