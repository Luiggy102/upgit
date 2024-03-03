package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/Luiggy102/upgit/cmd"
	e "github.com/Luiggy102/upgit/internal/extras"
	t "github.com/Luiggy102/upgit/types"
	"github.com/fatih/color"
)

var Paths = []string{}

// load saved paths
var pathsFile, _ = os.ReadFile(t.PathFile)
var _ = json.Unmarshal(pathsFile, &Paths)

func main() {
	// flags
	addPath := flag.String("add", " ", "add your git repo path")
	printStatus := flag.Bool("status", false, "print the paths status")
	pullChanges := flag.Bool("pull", false, "pull changes in saved repos")
	listRepos := flag.Bool("list", false, "list added repos")
	flag.Parse()

	var wg sync.WaitGroup

	// addPath flag
	switch *addPath {
	case " ":
		//
	default:
		path, err := filepath.Abs(*addPath)
		e.Check(err, "cannot convert string argument into a path")
		if e.IsGitPathValid(path) {
			if e.AlreadyAdded(Paths, path) {
				color.Yellow("Repo: '%s' already added\n", e.GetRepoName(path))
			} else {
				Paths = append(Paths, path)
				save()
				color.Green("Repo: '%s' added successfully\n", e.GetRepoName(path))
			}
			return
		} else {
			color.Red("not valid path/git repo, or don't have remote origin")
			return
		}
	}

	// print status flag
	if *printStatus {
		for _, path := range Paths {
			wg.Add(1)
			go cmd.PrintStatus(path, &wg)
		}
		wg.Wait()
		return
	}

	// pull changes flags
	if *pullChanges {
		okPaths, _ := cmd.CheckStatus(Paths)
		for _, path := range okPaths {
			wg.Add(1)
			cmd.PullChanges(path, &wg)
		}
		wg.Wait()
		return
	}

	// ListRepos flag
	if *listRepos {
		cmd.ListRepos(Paths)
		return
	}

	fmt.Println("use -h or --help for usage info")
}

// save path
func save() {
	err := os.MkdirAll(t.PathDir, 0700)
	e.Check(err, "can't create config dir")
	b, err := json.MarshalIndent(Paths, "", "\t")
	e.Check(err, "can't create json config file")
	err = os.WriteFile(t.PathFile, b, 0644)
	e.Check(err, "can't create config file")
}
