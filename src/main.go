package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", IndexRoute)
	mux.HandleFunc("/fact/", FactRoute)
	fmt.Println("Staring http-facter on port 4023.")
	http.ListenAndServe(":4023", mux)
}
