package main

import (
	"flag"
	"fmt"
	"os"
)

// constants
var DEVLOGCLI_VERSION = "v0.0.0"

// flags:
var versionFlag bool

func main() {

	// flag defaults
	const (
		defaultVersion = false
		usageVersion   = "Prints the version of Devlog-cli"
	)

	flag.BoolVar(&versionFlag, "version", defaultVersion, usageVersion)
	flag.BoolVar(&versionFlag, "v", defaultVersion, usageVersion+" (shorthand)")

	flag.Parse()

	if versionFlag {
		fmt.Print("devlog-cli ", DEVLOGCLI_VERSION, "\n")
		os.Exit(0)
	}
}
