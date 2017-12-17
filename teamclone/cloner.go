package teamclone

import (
	"github.com/steinfletcher/github-team-clone/github"
	"errors"
	"fmt"
	"os/exec"
	"sync"
)

type Cloner interface {
	CloneTeamRepos(org string, team string) error
}

type teamCloner struct {
	githubCli github.Github
	dir string
}

func NewCloner(g github.Github, dir string) Cloner {
	return &teamCloner{g, dir}
}

func (tC * teamCloner) CloneTeamRepos(org string, team string) error {
	err, teams := tC.githubCli.OrganisationTeams(org)
	if err != nil {
		return err
	}

	err, teamId := teamId(teams, team)
	if err != nil {
		return err
	}

	err, repos := tC.githubCli.TeamRepos(teamId)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	for _, repo := range repos {
		fmt.Println(fmt.Sprintf("Cloning %s", repo.Name))
		wg.Add(1)
		go clone(&wg, repo.SshUrl, repo.Name, tC.dir)
	}

	wg.Wait()
	return nil
}

func clone(wg *sync.WaitGroup, sshUrl string, repoName string, dir string) {
	defer wg.Done()

	out, err := exec.Command(
		"git", "clone", sshUrl, fmt.Sprintf("%s/%s", dir, repoName)).Output()

	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Printf("%s", out)
	}
}

func teamId(teams []github.Team, team string) (error, int) {
	for _, t := range teams {
		if t.Name == team {
			return nil, t.Id
		}
	}
	return errors.New(fmt.Sprintf("No team with name=%s exists with org", team)), 0
}