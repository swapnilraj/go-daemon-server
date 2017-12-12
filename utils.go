package main

import (
	"fmt"
	"os"
)

// exit exits the current program
func exit() {
	os.Exit(0)
}

// getVanityPrime generates the vanity prime
func getVanityPrime(vanity string) string {
	p, err := vanityPrime(vanity)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return p.Text(16) + "\n"
}
