package main

import (
	"net/http"
)

// Init starts the http server on the given port
func Init(port string) {
	// register route handler
	http.HandleFunc("/.well-known/vanityprime", primeHandler)
	http.HandleFunc("/.well-known/vpexit", exitHandler)

	// start listening
	http.ListenAndServe(port, nil)
}
