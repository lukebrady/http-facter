package main

import (
	"fmt"
	"net/http"
	"regexp"
)

//IndexRoute is a simple way to learn about http-facter.
func IndexRoute(w http.ResponseWriter, r *http.Request) {
	forbidden := "Error: Cannot access http-facter with " + r.Method + "."
	indexResponse := "HTTP-Facter\n"
	indexResponse += "-----------\n"
	indexResponse += "http-facter can be used to get facts from a remote server.\n\n"
	indexResponse += "For example, use /fact/os to gather facts about the OS.\n"
	// Check to make sure the API is not being accesses by POST, PUT, etc.
	if r.Method != "GET" {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(forbidden))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(indexResponse))
	}
}

// FactRoute is the handler that can handle all GET requests to Facter.
func FactRoute(w http.ResponseWriter, r *http.Request) {
	forbidden := "Error: Cannot access http-facter with " + r.Method + "."
	// Check to make sure the API is not being accesses by POST, PUT, etc.
	if r.Method != "GET" {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(forbidden))
	} else {
		// Translate the route into a command.
		command, err := TranslateRoute(r.RequestURI)
		if err != nil {

		}
		// Run facter on the system and return the output.
		output, err := RunFacter(command)
		if err != nil {
			fmt.Println("Error")
		}

		w.WriteHeader(http.StatusOK)
		w.Write(output)
	}
}

// TranslateRoute takes an http path and translates it into a call to Facter.
func TranslateRoute(route string) (string, error) {
	// Create a regular expression object to replace the /fact/ route.
	re1, err := regexp.Compile("/fact/")
	if err != nil {
		return "", err
	}
	// Replace the /fact/ route.
	reroute := re1.ReplaceAllString(route, "")
	// Create another regular expression object to replace all "/" with ".".
	re2, err := regexp.Compile("/")
	if err != nil {
		return "", err
	}
	// Replace all /.
	factCmd := re2.ReplaceAllString(reroute, ".")
	fmt.Println(factCmd)
	// Return the newly created command.
	return factCmd, nil
}
