package cli

import (
	"fmt"
	"strings"
	"time"

	"github.com/giuseppe-g-gelardi/git-sessionizer/api"

	"github.com/briandowns/spinner"
	"github.com/charmbracelet/log"
	"github.com/manifoldco/promptui"
)

var API_URL = "https://api.github.com/user/repos?page={PAGE}&per_page={PER_PAGE}&visibility=all"

type PartialRepo struct {
	Name        string `json:"name"`
	Http_url    string `json:"html_url"`
	Ssh_url     string `json:"ssh_url"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
}

func RepoSelection(token string) {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond) // Build our new spinner])
	s.Start()                                                    // Start the spinner
	s.Suffix = " Fetching all repos..."
	s.Color("cyan")
	repos, err := api.FetchAllUserRepos(API_URL, token)
	if err != nil {
		log.Errorf("Error: %v", err)
	}
	s.Stop() // Stop the spinner

	// convert repos to cli.PartialRepo and pass to cli.RepoPrompt to display
	var cliRepos []PartialRepo
	for _, repo := range repos {
		cliRepos = append(cliRepos, PartialRepo(repo))
	}
	repoPrompt(cliRepos)
}

func repoPrompt(repos []PartialRepo) {
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
		return
	}

	fmt.Printf("You choose number %d: %s\n", i+1, repos[i].Name)
}

// func Prompt_ui() {
// 	// Define a list of options
// 	items := []string{"Apple", "Banana", "Orange", "Grapes", "Strawberry"}
//
// 	// Create a Select prompt
// 	prompt := promptui.Select{
// 		Label: "Select a fruit",
// 		Items: items,
// 	}
//
// 	// Show the prompt to the user
// 	_, result, err := prompt.Run()
//
// 	if err != nil {
// 		fmt.Printf("Prompt failed %v\n", err)
// 		return
// 	}
//
// 	fmt.Printf("You selected: %s\n", result)
// }

//
// func RepoTest() {
//
// 	type pepper struct {
// 		Name     string
// 		HeatUnit int
// 		Peppers  int
// 	}
//
// 	peppers := []pepper{
// 		{Name: "Bell Pepper", HeatUnit: 0, Peppers: 0},
// 		{Name: "Banana Pepper", HeatUnit: 100, Peppers: 1},
// 		{Name: "Poblano", HeatUnit: 1000, Peppers: 2},
// 		{Name: "Jalapeño", HeatUnit: 3500, Peppers: 3},
// 		{Name: "Aleppo", HeatUnit: 10000, Peppers: 4},
// 		{Name: "Tabasco", HeatUnit: 30000, Peppers: 5},
// 		{Name: "Malagueta", HeatUnit: 50000, Peppers: 6},
// 		{Name: "Habanero", HeatUnit: 100000, Peppers: 7},
// 		{Name: "Red Savina Habanero", HeatUnit: 350000, Peppers: 8},
// 		{Name: "Dragon’s Breath", HeatUnit: 855000, Peppers: 9},
// 	}
//
// 	templates := &promptui.SelectTemplates{
// 		Label:    "   {{ . | faint }}?",
// 		Active:   "-> {{ .Name | blue }} ({{ .HeatUnit | red }})",
// 		Inactive: "   {{ .Name | cyan }} ({{ .HeatUnit | red }})",
// 		Selected: "   {{ .Name | red | cyan }}",
// 		Details: `
// --------- Pepper ----------
// {{ "Name:" | faint }}	{{ .Name }}
// {{ "Heat Unit:" | faint }}	{{ .HeatUnit }}
// {{ "Peppers:" | faint }}	{{ .Peppers }}`,
// 	}
//
// 	searcher := func(input string, index int) bool {
// 		pepper := peppers[index]
// 		name := strings.Replace(strings.ToLower(pepper.Name), " ", "", -1)
// 		input = strings.Replace(strings.ToLower(input), " ", "", -1)
//
// 		return strings.Contains(name, input)
// 	}
//
// 	prompt := promptui.Select{
// 		Label:     "Spicy Level",
// 		Items:     peppers,
// 		Templates: templates,
// 		Size:      4,
// 		Searcher:  searcher,
// 	}
//
// 	i, _, err := prompt.Run()
//
// 	if err != nil {
// 		fmt.Printf("Prompt failed %v\n", err)
// 		return
// 	}
//
// 	fmt.Printf("You choose number %d: %s\n", i+1, peppers[i].Name)
//
// }
