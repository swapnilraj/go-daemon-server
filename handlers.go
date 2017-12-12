package main

import (
	"fmt"
	"net/http"
)

// primeHandler responds with a vanity prime
func primeHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	vals := query["vs"]
	if len(vals) > 0 {
		fmt.Fprintf(w, getVanityPrime(vals[0]))
	} else {
		fmt.Fprintf(w, "no vs param")
	}
}

// exitHandler implements the exit program ability
func exitHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "exiting")
	go exit()
}
