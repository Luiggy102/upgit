package cmd

import (
	"fmt"
	"regexp"

	"github.com/Luiggy102/upgit/internal/extras"
	"github.com/fatih/color"
)

func RemoveRepo(repoName string, paths []string) []string {
	if len(paths) > 0 {
		for i := 0; i < len(paths); i++ {
			// match repo name with the path
			found, err := regexp.MatchString(repoName, paths[i])
			extras.Check(err, "")

			if found {
				fmt.Println("Deleted repo", paths[i])
				paths = append(paths[:i], paths[i+1:]...)
				break
			} else {
				fmt.Println("Coudn't find repo")
			}
		}
		return paths
	} else {
		color.Yellow("No repos added yet")
		return nil
	}
}
