# github-org-clone

[![CircleCI](https://circleci.com/gh/steinfletcher/github-org-clone.svg?style=svg&circle-token=063b1b1e0354cc424a2823c33ff4a2b66e029bae)](https://circleci.com/gh/steinfletcher/github-org-clone)

A simple cli app to clone all repos managed by a github organisation or team.
Requires that you pass a github api key (personal access token) and github username to the script or set the `GITHUB_TOKEN` and `GITHUB_USER` environment variable. See the help output below.

## Install

The following script will install a binary from a tagged release 

    curl https://raw.githubusercontent.com/steinfletcher/github-org-clone/master/download.sh | sh
    mv github-org-clone /usr/local/bin 

Or install from master using go

    go get github.com/steinfletcher/github-org-clone

## Use

Export env vars in `~/.bashrc` or equivalent

    export GITHUB_USER=<your github username>
    export GITHUB_TOKEN=<a github personal access token with clone repo privileges>

(Alternatively supply these as flags to the command `--username` and `--token`).

Clone team repos

    github-org-clone --org MyOrg --team MyTeam

Clone organisation repos

    github-org-clone -o MyOrg

Override the default location

    github-org-clone -o MyOrg -d ~/projects/work

View docs

    github-org-clone -h
