# Upgit

Cli programmg to check the status of multiple git repos that have a remote origin and make pulls on the repos that don't have nothing to commit, written in Go.  
Made for personal use, for linux and MacOs.

## Features

* Add many repos
* Check multiple status of git repos
* Make multiple pull to git repos
* Display a notification of the git status (linux only)
* Display a notification of the git pull log (linux only)

## Dependencies

* `dunst` (for linux notifications)

## Instalation

Use the go install command if you have go installed
```bash
go install github.com/Luiggy102/upgit@latest
```

## Usage

#### Show help
Use `-h` or `-help` for showing the options

```bash
upgit -h
```

#### Add a repo path
Use `-a` or adding a path to a git repo with a remote origin

```bash
upgit -a <<REPO PATH>>
```

#### List added repos
Use `-l` for show a list of added repos

```bash
upgit -l
```

#### Remove added repo
Use `-r` for delete added repo

```bash
upgit -r <<REPO NAME>>
```

#### Print status
Use `-s` for print the status of added repos

```bash
upgit -s
```

#### Make pull
Use `-ll` for make a pull of added repos

```bash
upgit -l
```

## Notifications (linux only)

#### Print status
Use `-sn` for display a notification with the status of repos

```bash
upgit -sn
```

#### Pulls
Use `-lln` for display a notification with the pulls progress

```bash
upgit -lln
```
