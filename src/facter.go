package main

import (
	"os/exec"
)

// RunFacter executes facter with the command given.
func RunFacter(command string) (string, error) {
	facter := exec.Command("facter", command)
	err := facter.Run()
	if err != nil {
		return "", err
	}
	output, err := facter.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// FacterToJSON converts facter's output to JSON and then
// returns the JSON so that it can be served.
func FacterToJSON(output string) []byte {
	re := 
}
