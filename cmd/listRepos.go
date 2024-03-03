package cmd

import (
	"fmt"

	"github.com/Luiggy102/upgit/internal/extras"
	"github.com/fatih/color"
)

func ListRepos(paths []string) {
	if len(paths) == 0 {
		color.Yellow("No repos added yet")
	} else {
		for _, path := range paths {
			fmt.Printf("* Repo: %s\n  Path: %s \n",
				color.GreenString(extras.GetRepoName(path)),
				color.MagentaString(path))
		}
	}

}
