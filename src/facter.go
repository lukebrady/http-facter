package main

import (
	"os/exec"
	"regexp"
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
	// Take out all double quotes from the facter output.
	re1, err := regexp.Compile("\"")
	if err != nil {

	}
	// Replace the double quotes with an empty string.
	quoteFacter1 := re1.ReplaceAllString(output, "")
	// Find all string and numeric values in the facter output.
	re2, err := regexp.Compile("([A-Z,a-z,0-100]+)")
	if err != nil {

	}
	// Add double quotes to every value in the new JSON string.
	quoteFacter2 := re2.ReplaceAllString(quoteFacter1, "\"\"")
	// Replace the => with :.
	re3, err := regexp.Compile("=>")
	if err != nil {

	}
	quoteFacter3 := re3.ReplaceAllString(quoteFacter2, ":")

	return []byte(quoteFacter3)
}
