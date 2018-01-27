package cloner

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/steinfletcher/github-team-clone/github"
)

func TestGetTeamId(t *testing.T) {
	teams := []github.Team{
		{Name:"team1", Id:1},
		{Name:"team2", Id:51},
		{Name:"team3", Id:101},
	}

	_, id := teamId(teams, "team3")

	assert.Equal(t, id, 101)
}

func TestGetTeamIdErrorsIfNotFound(t *testing.T) {
	teams := []github.Team{
		{Name:"team1", Id:1},
	}

	err, _ := teamId(teams, "team3")

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "No team with name")
}