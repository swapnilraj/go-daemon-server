package main

import (
	"os"
)

const defaultPort = ":8081"

// process starts the vanity prime server on either the default port
// or the one provided in the args
func process() {
	port := defaultPort
	if len(os.Args) > 1 {
		port = ":" + os.Args[1]
	}
	Init(port)
}

func main() {
	Run(process)
}
