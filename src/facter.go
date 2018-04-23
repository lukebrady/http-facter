package main

import (
	"os/exec"
	"strings"
	"regexp"
)

// GetFactFromURI takes a URI string and translates it into a format
// facter can use. If the request was /os/family, GetFactFromURI would
// return os.family so that it can be used by factor.
func GetFactFromURI(uri string) (string, error) {
	var factPath string
	// Create a regular expression object
	re := regexp.MustCompile(".")
	// Replace all forward slashes with periods
	for x, char := range uri {
		if x != 0 {
			strings.Replace(uri,"/",".",0)
		}
	}
	return uri, nil
}

// RunFacter fetches the fact from the server and returns it's value.
func RunFacter(fact string) ([]byte, error) {
	var facter *exec.Cmd
	// Create the facter command.
	facter = exec.Command("facter", fact)
	// Run the facter command and get its output.
	facter.Run()
}