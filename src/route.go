package main

import "net/http"
import "fmt"

//IndexRoute is a simple way to learn about http-facter.
func IndexRoute(w http.ResponseWriter, r *http.Request) {
	forbidden := "Error: Cannot access http-facter with " + r.Method + "."
	indexResponse := "HTTP-Facter\n"
	indexResponse += "-----------\n"
	indexResponse += "http-facter can be used to get facts from a remote server.\n\n"
	indexResponse += "Use /os to gather facts about the OS.\n"
	indexResponse += "Use /memory to gather facts about the systems memory.\n"
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
		TranslateRoute(r.RequestURI)
	}
}

// TranslateRoute takes an http path and translates it into a call to Facter.
func TranslateRoute(route string) {
	fmt.Println(route)
}
