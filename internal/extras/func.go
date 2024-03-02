package extras

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
)

// Function to check errors fast
func Check(err error, cause string) {
	if err != nil {
		fmt.Println(cause)
		panic(err)
	}
}

// Function to check if a command exits
func CommandExist(command string) bool {
	_, err := exec.LookPath(command)
	if err != nil {
		return false
	} else {
		return true
	}
}

// Function to check if path has a git repo
func IsGitPathValid(path string) bool {

	// check if path is valid
	_, err := os.Stat(path)
	if err != nil {
		log.Fatal("Path not valid")
		return false
	}

	// check if had a git repo
	gitPath := fmt.Sprintf("%s/.git", path)
	_, err = os.Stat(gitPath)
	if err != nil {
		log.Fatal("Path with no git repo")
		return false
	}

	// check if had remote
	gitPath = fmt.Sprintf("%s/config", gitPath)
	if CommandExist("cat") {
		// check the output of the config file
		cmd := fmt.Sprintf("cat %s", gitPath)
		out, err := exec.Command("/bin/sh", "-c", cmd).Output()
		Check(err, "cant read git config file")

		found, err := regexp.MatchString("remote", string(out))
		Check(err, "git repo don't have remote")

		if found {
			return true
		} else {
			return false
		}

	} else {
		log.Fatal("Cat dependency not fount, can't check repo")
		return false
	}

}
