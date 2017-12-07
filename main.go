package main

import (
	"os"

	"github.com/swapnilraj/go-daemon-server/daemon"
	"github.com/swapnilraj/go-daemon-server/server"
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
