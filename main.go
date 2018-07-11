package main

import (
	"fmt"
	"github.com/steinfletcher/github-org-clone/cloner"
	"github.com/steinfletcher/github-org-clone/github"
	"github.com/steinfletcher/github-org-clone/shell"
	"github.com/urfave/cli"
	"log"
	"os"
	"time"
)

var (
	version = "dev"
	commit  = ""
	date    = time.Now().String()
)

func main() {
	app := cli.NewApp()
	app.Author = "Stein Fletcher"
	app.Name = "github-org-clone"
	app.Usage = "clone github team repos"
	app.UsageText = "github-org-clone -o MyOrg -t MyTeam"
	app.Version = version
	app.EnableBashCompletion = true
	app.Description = "A simple cli to clone all the repos managed by a github team"
	app.Metadata = map[string]interface{}{
		"commit": commit,
		"date":   date,
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "org, o",
			Usage: "github organisation",
		},
		cli.StringFlag{
			Name:  "team, t",
			Usage: "github team",
		},
		cli.StringFlag{
			Name:   "username, u",
			Usage:  "github username",
			EnvVar: "GITHUB_USER,GITHUB_USERNAME",
		},
		cli.StringFlag{
			Name:   "token, k",
			Usage:  "github personal access token",
			EnvVar: "GITHUB_TOKEN,GITHUB_API_KEY,GITHUB_PERSONAL_ACCESS_TOKEN",
		},
		cli.StringFlag{
			Name:  "dir, d",
			Usage: "directory to clone into. Defaults to the org name or org/team name if defined",
		},
	}

	app.Action = func(c *cli.Context) error {
		username := c.String("username")
		token := c.String("token")
		team := c.String("team")
		org := c.String("org")
		dir := c.String("dir")

		if len(username) == 0 {
			die("env var GITHUB_USERNAME or flag -u must be set", c)
		}

		if len(token) == 0 {
			die("env var GITHUB_TOKEN or flag -k must be set", c)
		}

		if len(org) == 0 {
			die("github organisation (-o) not set", c)
		}

		if len(dir) == 0 {
			if len(team) == 0 {
				dir = org
			} else {
				if _, err := os.Stat(org); os.IsNotExist(err) {
					os.Mkdir(org, os.ModePerm)
				}
				dir = fmt.Sprintf("%s/%s", org, team)
			}
		}

		sh := shell.NewShell()
		githubCli := github.NewGithub(username, token)
		cl := cloner.NewCloner(githubCli, sh, dir)

		err := cl.Clone(org, team)
		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		}

		return nil
	}

	app.Run(os.Args)
}

func die(msg string, c *cli.Context) {
	cli.ShowAppHelp(c)
	log.Fatal(msg)
}
