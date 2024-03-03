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
	addPath := flag.String("a", " ", "add your git repo path")
	printStatus := flag.Bool("s", false, "print the paths status")
	pullChanges := flag.Bool("ll", false, "pull changes in saved repos")
	listRepos := flag.Bool("l", false, "list added repos")
	removeRepo := flag.String("r", " ", "remove added repo with the repo name")

	statusNotification := flag.Bool("sn", false, "print a notificacion with the status")
	pullNotificatin := flag.Bool("lln", false, "print a notificacion with the pull log")

	// flag.Usage = func() {
	// 	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	//
	// 	flag.VisitAll(func(f *flag.Flag) {
	// 		fmt.Fprintf(os.Stderr, "-%s    \t%s\n", f.Name, f.Usage)
	// 	})
	// }

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
		if len(Paths) > 0 {
			for _, path := range Paths {
				wg.Add(1)
				go cmd.PrintStatus(path, &wg)
			}
			wg.Wait()
		} else {
			color.Yellow("No repos added yet")
		}
		return
	}

	// pull changes flags
	if *pullChanges {
		if len(Paths) > 0 {
			okPaths, _ := cmd.CheckStatus(Paths)
			for _, path := range okPaths {
				wg.Add(1)
				cmd.PullChanges(path, &wg)
			}
			wg.Wait()
		} else {
			color.Yellow("No repos added yet")
		}
		return
	}

	// ListRepos flag
	if *listRepos {
		cmd.ListRepos(Paths)
		return
	}

	// removeRepo flag
	switch *removeRepo {
	case " ":
		//
	default:
		Paths = cmd.RemoveRepo(*removeRepo, Paths)
		save()
		return
	}

	// status notificacion flag
	if *statusNotification {
		if len(Paths) > 0 {
			for _, path := range Paths {
				wg.Add(1)
				go cmd.StatusNotification(path, &wg)
			}
			wg.Wait()
		} else {
			color.Yellow("No repos added yet")
		}
		return
	}

	// pull notificacion flag
	if *pullNotificatin {
		if len(Paths) > 0 {
			okPaths, _ := cmd.CheckStatus(Paths)
			for i, path := range okPaths {
				cmd.PullNotification(path, i+1, len(okPaths))
			}
		} else {
			color.Yellow("No repos added yet")
		}
		return
	}

	fmt.Println(`Be aware of multiple git repos and make multiple pulls!
	        use -h or --help for usage info`)
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
