package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	// IndexRoute handles traffic to the route uri. This route
	// writes out some instructions about the command.
	mux.HandleFunc("/", IndexRoute)
	// FactRoute handles all fact requests and sends the command
	// to facter. After facter runs the output is translated to JSON
	// and is then returned.
	mux.HandleFunc("/fact/", FactRoute)
	fmt.Println("Staring http-facter on port 4400.")
	http.ListenAndServe(":4400", mux)
}
