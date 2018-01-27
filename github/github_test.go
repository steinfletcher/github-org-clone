package github

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestFetchTeams(t *testing.T) {
	defer gock.Off()

	gock.New("https://api.github.com").
		MatchHeader("Authorization", "Basic dXNlcm5hbWU6cGFzc3dvcmQ=").
		Get("/orgs/MyOrg/teams").
		Reply(200).
		File("../testdata/teamsResponse.json")

	githubCli := NewGithub("username", "password")

	_, teams := githubCli.Teams("MyOrg")

	assert.Equal(t, teams[0].Name, "Winners")
	assert.Equal(t, teams[0].Id, 2285788)
}

func TestFetchTeamRepos(t *testing.T) {
	defer gock.Off()

	teamReposResp("1", "<https://api.github.com/teams/2285789/repos?page=2>; rel=\"next\", <https://api.github.com/teams/2285789/repos?page=3>; rel=\"last\"", "teamReposResponsePage1.json")
	teamReposResp("2", "<https://api.github.com/teams/2285789/repos?page=1>; rel=\"prev\", <https://api.github.com/teams/2285789/repos?page=3>; rel=\"next\", <https://api.github.com/teams/2285789/repos?page=1>; rel=\"first\", <https://api.github.com/teams/2285789/repos?page=3>; rel=\"last\"", "teamReposResponsePage2.json")
	teamReposResp("3", "<https://api.github.com/teams/2285789/repos?page=2>; rel=\"prev\", <https://api.github.com/teams/2285789/repos?page=1>; rel=\"first\"", "teamReposResponsePage3.json")

	githubCli := NewGithub("username", "password")

	_, teams := githubCli.TeamRepos(2285789)

	var repos []string
	for _, team := range teams {
		repos = append(repos, team.Name)
	}

	assert.Contains(t, repos, "some-repo1")
	assert.Contains(t, repos, "some-repo2")
	assert.Contains(t, repos, "some-repo3")
}

func teamReposResp(page string, linkHeader string, file string) {
	gock.New("https://api.github.com").
		MatchHeader("Authorization", "Basic dXNlcm5hbWU6cGFzc3dvcmQ=").
		Get("/teams/2285789/repos").
		MatchParam("page", page).
		Reply(200).
		SetHeader("Link", linkHeader).
		File("../testdata/" + file)
}