package cmd

import (
	"fmt"
	"os/exec"
	"regexp"
	"runtime"
	"sync"

	e "github.com/Luiggy102/upgit/internal/extras"
	"github.com/fatih/color"
)

func StatusNotification(path string, wg *sync.WaitGroup) {
	defer wg.Done()

	if runtime.GOOS == "darwin" {

		color.Red("Feature not available for OsX")

	} else {

		if e.CommandExist("dunst") && e.CommandExist("dunstify") {

			cmd := fmt.Sprintf("cd %s && git status", path)
			out, _ := exec.Command("/bin/sh", "-c", cmd).Output()
			status := string(out)
			found, _ := regexp.MatchString("nothing to commit", status)

			if found {
				cmd = fmt.Sprintf("dunstify '  Ruta: %s' 'Status: Ok'",
					path)
				exec.Command("/bin/sh", "-c", cmd).Run()
			} else {
				cmd = fmt.Sprintf("dunstify '  Ruta: %s' 'Status: %s' -u critical -t 15000",
					path, status)
				exec.Command("/bin/sh", "-c", cmd).Run()
			}

		} else {

			color.Red("Don't have dunst dependency installed")

		}

	}

}
