package cloner

import (
	"errors"
	"fmt"
	"sync"

	"github.com/steinfletcher/github-org-clone/github"
	"github.com/steinfletcher/github-org-clone/shell"
)

type Cloner interface {
	Clone(org string, team string) error
}

type teamCloner struct {
	githubCli github.Github
	shell     shell.Shell
	dir       string
}

func NewCloner(g github.Github, shell shell.Shell, dir string) Cloner {
	return &teamCloner{g, shell, dir}
}

func (tC *teamCloner) Clone(org string, team string) error {
	var repos []github.Repo
	var err error

	if team == "" {
		err, repos = tC.githubCli.OrgRepos(org)
		if err != nil {
			return err
		}
	} else {
		e, teams := tC.githubCli.Teams(org)
		if e != nil {
			return e
		}

		e, teamId := teamId(teams, team)
		if e != nil {
			return e
		}

		e, repos = tC.githubCli.TeamRepos(teamId)
		if e != nil {
			return e
		}
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
