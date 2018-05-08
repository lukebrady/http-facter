package main

import (
	"os/exec"
)

// RunFacter executes facter with the command given.
func RunFacter(command string) ([]byte, error) {
	// Execute the facter command with JSON output.
	facter := exec.Command("facter", "-j", command)
	// Get the output of the command and return err if error occurs.
	output, err := facter.Output()
	if err != nil {
		return nil, err
	}
	// Return the facter output as a byte slice.
	return output, nil
}
