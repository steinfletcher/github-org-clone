# github-org-clone

[![Build Status](https://travis-ci.org/steinfletcher/github-org-clone.svg?branch=master)](https://travis-ci.org/steinfletcher/github-org-clone)

A simple cli app to clone all repos managed by a github organisation or team.
Requires that you pass a github api key (personal access token) and github username to the script or set the `GITHUB_TOKEN` and `GITHUB_USER` environment variable. See the help output below.

## Install

The following script will install a binary from a tagged release 

```bash
curl https://raw.githubusercontent.com/steinfletcher/github-org-clone/master/download.sh | sh
mv github-org-clone /usr/local/bin
```

Or install from master using go

```bash
go get github.com/steinfletcher/github-org-clone
```

## Use

Export env vars in `~/.bashrc` or equivalent

```bash
export GITHUB_USER=<your github username>
export GITHUB_TOKEN=<a github personal access token with clone repo privileges>
```

(Alternatively supply these as flags to the command `--username` and `--token`).

Clone team repos

```bash
github-org-clone --org MyOrg --team MyTeam
```

Clone organisation repos

```bash
github-org-clone -o MyOrg
```

Override the default location

```bash
github-org-clone -o MyOrg -d ~/projects/work
```

Override the github api url

```bash
github-org-clone -o MyOrg -a https://mycustomdomain.com
```

View docs

```bash
github-org-clone -h
```
