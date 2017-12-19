package teamclone

import (
	"github.com/steinfletcher/github-team-clone/github"
	"errors"
	"fmt"
	"sync"
	"github.com/steinfletcher/github-team-clone/shell"
)

type Cloner interface {
	CloneTeamRepos(org string, team string) error
}

type teamCloner struct {
	githubCli github.Github
	shell shell.Shell
	dir string
}

func NewCloner(g github.Github, shell shell.Shell, dir string) Cloner {
	return &teamCloner{g, shell, dir}
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
		go tC.clone(&wg, repo.SshUrl, repo.Name, tC.dir)
	}

	wg.Wait()
	return nil
}

func (tC *teamCloner) clone(wg *sync.WaitGroup, sshUrl string, repoName string, dir string) {
	defer wg.Done()
	tC.shell.Exec("git", []string{"clone", sshUrl, fmt.Sprintf("%s/%s", dir, repoName)})
}

func teamId(teams []github.Team, team string) (error, int) {
	for _, t := range teams {
		if t.Name == team {
			return nil, t.Id
		}
	}
	return errors.New(fmt.Sprintf("No team with name=%s exists with org", team)), 0
}