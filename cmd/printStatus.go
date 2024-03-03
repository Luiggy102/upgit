package cmd

import (
	"fmt"
	"os/exec"
	"regexp"
	"sync"

	"github.com/Luiggy102/upgit/internal/extras"
	"github.com/fatih/color"
)

func PrintStatus(path string, wg *sync.WaitGroup) {

	defer wg.Done()

	cmd := fmt.Sprintf("cd %s && git status", path)
	out, _ := exec.Command("/bin/sh", "-c", cmd).Output()
	status := string(out)

	found, _ := regexp.MatchString("nothing to commit", status)

	repoName := extras.GetRepoName(path)

	if found {
		status = fmt.Sprintf("Repo: %s\n%s",
			color.MagentaString(repoName),
			color.GreenString("Status: ok"))
		fmt.Println(status)
	} else {
		status = fmt.Sprintf("Repo: %s\nStatus: %s",
			color.MagentaString(repoName),
			color.RedString(status))
		fmt.Println(status)
	}

}
