package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	e "github.com/Luiggy102/upgit/internal/extras"
	t "github.com/Luiggy102/upgit/types"
)

var Paths = []string{}

// load paths
var pathsFile, _ = os.ReadFile(t.PathFile)
var _ = json.Unmarshal(pathsFile, &Paths)

func main() {
	// flags
	addPath := flag.String("a", "", "add a path to your git repo")
	flag.Parse()

	// addPath flag
	path, _ := filepath.Abs(*addPath)
	if e.IsGitPathValid(path) {
		if e.AlreadyAdded(Paths, path) {
			fmt.Println("Path already added")
		} else {
			Paths = append(Paths, path)
			save()
			fmt.Println("Path added successfully")
		}
	} else {
		fmt.Println("not valid path, not git repo, or no have remote")
	}
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
