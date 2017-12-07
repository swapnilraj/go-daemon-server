package main

import (
	"os"

	"github.com/swapnilraj/assignment/daemon"
	"github.com/swapnilraj/assignment/server"
)

const defaultPort = ":8081"

// process starts the vanity prime server on either the default port
// or the one provided in the args
func process() {
	port := defaultPort
	if len(os.Args) > 1 {
		port = ":" + os.Args[1]
	}
	server.Init(port)
}

func main() {
	daemon.Run(process)
}
