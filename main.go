package main

import (
	"os"

	"github.com/ta9mi1shi1/toggl/cmd"
)

type exitCode int

const (
	exitOK    exitCode = 0
	exitError exitCode = 1
)

func main() {
	os.Exit(int(run()))
}

func run() exitCode {
	rootCmd := cmd.NewCmdRoot()
	if err := rootCmd.Execute(); err != nil {
		return exitError
	}
	return exitOK
}
