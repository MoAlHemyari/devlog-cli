package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	versionArgValue bool
)

const CURRENT_VERSION string = "v0.0.0"

func main() {
	const (
		versionDefault bool   = false
		versionUsage   string = "This returns the version of devlog-cli"
	)

	flag.BoolVar(&versionArgValue, "version", versionDefault, versionUsage)
	flag.BoolVar(&versionArgValue, "v", versionDefault, versionUsage)

	flag.Parse()

	if versionArgValue {
		fmt.Print("devlog-cli "+CURRENT_VERSION, "\n")
		os.Exit(0)
	}
}
