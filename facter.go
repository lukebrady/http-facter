package main

import (
	"log"
	"os"
	"os/exec"
)

// RunFacter executes facter with the command given.
func RunFacter(command string) ([]byte, error) {
	// Open the http-facter log file.
	lFile, err := os.OpenFile("/var/log/httpfacter.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	defer lFile.Close()
	// Set the log output to lFile.
	log.SetOutput(lFile)
	// Append to the log that a new request has come in and facter is about to run.
	log.Printf("Facter is looking up %s.\n", command)
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
