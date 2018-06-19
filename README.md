# github-org-clone

[![CircleCI](https://circleci.com/gh/steinfletcher/github-org-clone.svg?style=svg&circle-token=063b1b1e0354cc424a2823c33ff4a2b66e029bae)](https://circleci.com/gh/steinfletcher/github-org-clone)

A simple cli app to clone all repos managed by a github organisation or team.
Requires that you pass a github api key (personal access token) and github username to the script or set the `GITHUB_TOKEN` and `GITHUB_USER` environment variable. See the help output below.

## Install

The following script will install a binary from a tagged release 

    curl https://raw.githubusercontent.com/steinfletcher/github-org-clone/master/install.sh | sh

Or install from master using go

    go get github.com/steinfletcher/github-org-clone

## Use

Clone team repos

    github-org-clone -o MyOrg -t MyTeam


Clone organisation repos

    github-org-clone -o MyOrg


