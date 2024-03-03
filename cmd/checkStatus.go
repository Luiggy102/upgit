package cmd

import (
	"fmt"
	"os/exec"
	"regexp"
)

func CheckStatus(paths []string) (okPaths []string, notOkPaths []string) {

	for _, path := range paths {

		cmd := fmt.Sprintf("cd %s && git status", path)
		out, _ := exec.Command("/bin/sh", "-c", cmd).Output()
		status := string(out)

		found, _ := regexp.MatchString("nothing to commit", status)

		if found {
			okPaths = append(okPaths, path)
		} else {
			notOkPaths = append(notOkPaths, path)
		}
	}
	return
}
