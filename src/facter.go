package main

import (
	"os/exec"
)

// RunFacter executes facter with the command given.
func RunFacter(command string) ([]byte, error) {
	facter := exec.Command("facter", "-j", command)
	err := facter.Run()
	if err != nil {
		return nil, err
	}
	output, err := facter.Output()
	if err != nil {
		return nil, err
	}
	return []byte(output), nil
}
