package cmd

import (
	"fmt"
	"os/exec"
	"sync"

	"github.com/Luiggy102/upgit/internal/extras"
	"github.com/fatih/color"
)

func PullChanges(path string, wg *sync.WaitGroup) {

	defer wg.Done()

	cmd := fmt.Sprintf("cd %s && git pull origin main", path)
	out, _ := exec.Command("/bin/sh", "-c", cmd).Output()
	pullLog := string(out)

	repoName := extras.GetRepoName(path)

	fmt.Printf("Repo: %s\nLog: %s",
		color.MagentaString(repoName), color.GreenString(pullLog))

}
