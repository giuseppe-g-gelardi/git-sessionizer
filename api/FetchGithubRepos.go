package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type PartialRepo struct {
	Name        string `json:"name"`
	Http_url    string `json:"html_url"`
	Ssh_url     string `json:"ssh_url"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
}

func FetchAllUserRepos(url string, token string) ([]PartialRepo, error) {
	perPage := 100
	page := 1
	allRepos := []PartialRepo{}
	moreRepos := true

	for moreRepos {
		uri := strings.Replace(url, "{PAGE}", strconv.Itoa(page), -1)
		uri = strings.Replace(uri, "{PER_PAGE}", strconv.Itoa(perPage), -1)

		repos, err := fetchRepos(token, uri)
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

	// fmt.Printf("Fetched %d repositories!\n", len(allRepos))
	return allRepos, nil
}

func fetchRepos(token, url string) ([]PartialRepo, error) {
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
