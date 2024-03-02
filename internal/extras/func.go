package extras

import (
	"fmt"
	"os/exec"
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
