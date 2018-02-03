# github-org-clone

[![CircleCI](https://circleci.com/gh/steinfletcher/github-org-clone.svg?style=svg&circle-token=063b1b1e0354cc424a2823c33ff4a2b66e029bae)](https://circleci.com/gh/steinfletcher/github-org-clone)

A simple cli app to clone all repos managed by a github organisation or team.
Requires that you pass a github api key (personal access token) and github username to the script or set the `GITHUB_TOKEN` and `GITHUB_USER` environment variable. See the help output below.

## Install

    go get github.com/steinfletcher/github-org-clone

## Use

Clone team repos

    github-org-clone -o MyOrg -t MyTeam


Clone orgnisation repos

    github-org-clone -o MyOrg

```bash
NAME:
github-org-clone - clone github team repos

USAGE:
     $ github-org-clone -o MyOrg -t MyTeam

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
