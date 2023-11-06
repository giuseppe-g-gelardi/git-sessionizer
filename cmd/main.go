package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	// "io"
	"net/http"

	// "github.com/giuseppe-g-gelardi/git-sessionizer/auth"
	"github.com/charmbracelet/log"
	"github.com/giuseppe-g-gelardi/git-sessionizer/cli"
	c "github.com/giuseppe-g-gelardi/git-sessionizer/config"
	// "github.com/manifoldco/promptui"
)

type PartialRepo struct {
	Name        string `json:"name"`
	Http_url    string `json:"html_url"`
	Ssh_url     string `json:"ssh_url"`
	Description string `json:"description"`
}

var API_URL = "https://api.github.com/user/repos?page={PAGE}&per_page={PER_PAGE}&visibility=all"

func main() {
	// ! isAuth, err := auth.Authenticate() // auth.Authenticate()

	// get the .configrc rile
	conf, err := c.NewConfigManager().GetConfig(1)
	if err != nil {
		log.Errorf("Error: %v", err)
	}

	// get all user repos
    fmt.Println("Fetching all repos...")
	repos, err := AppendAllRepos(API_URL, conf.AccessToken)
	if err != nil {
		log.Errorf("Error: %v", err)
	}
    fmt.Println("Done!")

	// convert repos to cli.PartialRepo and pass to cli.RepoPrompt to display
	var cliRepos []cli.PartialRepo
	for _, repo := range repos {
		cliRepos = append(cliRepos, cli.PartialRepo(repo))
	}

	// prompt user to select a repo
	cli.RepoPrompt(cliRepos)
}

// ! ===================================================================== ! //

func AppendAllRepos(url string, token string) ([]PartialRepo, error) {
	perPage := 100
	page := 1
	allRepos := []PartialRepo{}
	moreRepos := true

	for moreRepos {
		uri := strings.Replace(url, "{PAGE}", strconv.Itoa(page), -1)
		uri = strings.Replace(uri, "{PER_PAGE}", strconv.Itoa(perPage), -1)

		repos, err := FetchGithubRepos(token, uri)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return nil, err
		}

		if len(repos) == 0 {
			moreRepos = false
			break
		} else {
			allRepos = append(allRepos, repos...)
			page++
		}
	}

	fmt.Printf("Fetched %d repositories!\n", len(allRepos))
	return allRepos, nil
}

func FetchGithubRepos(token, url string) ([]PartialRepo, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "token "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	var repos []PartialRepo
	err = json.NewDecoder(resp.Body).Decode(&repos)
	if err != nil {
		return nil, err
	}

	return repos, nil
}

// ! ===================================================================== ! //
// ! ===================================================================== ! //
// ! ===================================================================== ! //

// repos, err := FetchGithubRepos(API)
// if err != nil {
// 	log.Errorf("Error: %v", err)
// }
// log.Infof("Repos: %v", repos)

// func FetchGithubRepos(url string) ([]PartialRepo, error) {
// 	conf, err := c.NewConfigManager().GetConfig(1)
// 	if err != nil {
// 		log.Errorf("Error: %v", err)
// 	}

// 	client := &http.Client{}
// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		log.Errorf("Error: %v", err)
// 		return nil, err
// 	}

// 	req.Header.Add("Authorization", "token "+conf.AccessToken)

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Errorf("Error: %v", err)
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		log.Errorf("Error: %v", resp.Status)
// 		return nil, err
// 	}

// 	var repos []PartialRepo

// 	decoder := json.NewDecoder(resp.Body)
// 	if err := decoder.Decode(&repos); err != nil {
// 		log.Errorf("Error: %v", err)
// 		return nil, err
// 	}

// 	return repos, nil
// }
