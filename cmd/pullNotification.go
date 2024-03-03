package cmd

import (
	"fmt"
	"os/exec"
	"runtime"

	e "github.com/Luiggy102/upgit/internal/extras"
	"github.com/fatih/color"
)

func PullNotification(path string, id int, total int) {

	if runtime.GOOS == "darwin" {

		color.Red("Feature not available for OsX")

	}

	if e.CommandExist("dunst") && e.CommandExist("dunstify") {

		cmd := fmt.Sprintf("cd %s && git pull origin main", path)
		out, _ := exec.Command("/bin/sh", "-c", cmd).Output()
		pullLog := string(out)

		percentage := float64(id) / float64(total) * 100

		repoName := e.GetRepoName(path)

		// -r is the notificacion id, fot not stacking notitication
		cmd = fmt.Sprintf("dunstify '(%d/%d)   Finishing git pull: %s' '%s' -h int:value:%v -r 1234",
			id, total, repoName, pullLog, percentage)
		exec.Command("/bin/sh", "-c", cmd).Run()

	}

}
