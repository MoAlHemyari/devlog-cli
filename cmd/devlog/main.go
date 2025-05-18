package main

import (
	"fmt"
	"os"
)

var (
	showVersion bool
	showHelp    bool
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Print("ok. funny.")
		os.Exit(0)
	}

	for i := 0; i <= len(args); i++ {
		switch args[0] {
		case "-v", "--version":
			showVersion = true

		case "-h", "--help":
			showHelp = true

		default:
			fmt.Fprintf(os.Stderr, "Usage: %s [--version/-v] [--help/-h]\n", os.Args[0])
			os.Exit(2)
		}
	}

	if showVersion {
		fmt.Print("devlog-cli v0.0.0\n")
		os.Exit(0)
	}

	if showHelp {
		fmt.Print("devlog-cli is a developer tool that can be used to make participating in projects easier.\n")
		os.Exit(0)
	}

	os.Exit(0)
}
