# github-team-clone

A simple cli app to clone all repos managed by a github team

## Install

    go install github.com/steinfletcher/github-team-clone

## Usage

    github-team-clone -o MyOrg -t MyTeam

```bash
NAME:
github-team-clone - clone github team repos

USAGE:
     $ github-team-clone -o MyOrg -t MyTeam

VERSION:
0.0.1

DESCRIPTION:
A simple cli to clone all the repos managed by a github team

AUTHOR:
Stein Fletcher

COMMANDS:
help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
--org value, -o value       github organisation
--team value, -t value      github team
--username value, -u value  github username [$GITHUB_USER, $GITHUB_USERNAME]
--token value, -k value     github personal access token [$GITHUB_TOKEN, $GITHUB_API_KEY, $GITHUB_PERSONAL_ACCESS_TOKEN]
--dir value, -d value       directory to clone into (default: "src")
--help, -h                  show help
--version, -v               print the version
```