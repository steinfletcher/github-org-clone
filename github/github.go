package github

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"time"
	"encoding/json"
	"strings"
	"errors"
)

func NewGithub(username string, apiToken string) Github {
	return &githubCli{username, apiToken}
}

type githubCli struct {
	username string
	apiToken string
}

type Github interface {
	Teams(org string) (error, []Team)
	TeamRepos(teamId int) (error, []Repo)
	OrgRepos(org string) (error, []Repo)
}

type Team struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Repo struct {
	SshUrl string `json:"ssh_url"`
	Name   string `json:"name"`
}

func (g *githubCli) Teams(org string) (error, []Team) {
	err, resp := doGet(fmt.Sprintf("https://api.github.com/orgs/%s/teams", org), g.username, g.apiToken)
	defer resp.Body.Close()
	if err != nil {
		return err, nil
	}
	teams := make([]Team, 0)
	bytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bytes, &teams)
	return nil, teams
}

func (g *githubCli) TeamRepos(teamId int) (error, []Repo) {
	var page = 1
	var repos []Repo
	for {
		err, res := doGet(fmt.Sprintf("https://api.github.com/teams/%d/repos?page=%d", teamId, page), g.username, g.apiToken)
		if err != nil {
			return err, nil
		}

		bytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err, nil
		}
		result := make([]Repo, 0)
		json.Unmarshal(bytes, &result)
		repos = append(repos, result...)

		if hasNextPage(res) {
			page++
		} else {
			break
		}
	}
	return nil, repos
}

func (g *githubCli) OrgRepos(org string) (error, []Repo) {
	var page = 1
	var repos []Repo
	for {
		err, res := doGet(fmt.Sprintf("https://api.github.com/orgs/%s/repos?page=%d", org, page), g.username, g.apiToken)
		if err != nil {
			return err, nil
		}

		bytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err, nil
		}
		result := make([]Repo, 0)
		json.Unmarshal(bytes, &result)
		repos = append(repos, result...)

		if hasNextPage(res) {
			page++
		} else {
			break
		}
	}
	return nil, repos
}

func doGet(url string, username string, apiToken string) (error, *http.Response) {
	httpClient := http.Client{Timeout: time.Second * 5}
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf(url), nil)
	request.SetBasicAuth(username, apiToken)

	response, err := httpClient.Do(request)

	if err != nil {
		return err, nil
	}

	if !isSuccess(response) {
		return errors.New(fmt.Sprintf("Error, github returned status=%s", response.StatusCode)), nil
	}

	return nil, response
}

func isSuccess(r *http.Response) bool {
	return r.StatusCode >= 200 && r.StatusCode < 300
}

func hasNextPage(response *http.Response) bool {
	h := response.Header.Get("link")
	pageLinks := strings.Split(h, ",")
	for _, link := range pageLinks {
		if strings.Contains(link, "rel=\"next\"") {
			return true
		}
	}
	return false
}
