package main

import (
	"os"

	"github.com/instruqt/aws-nuke/v3/cmd"
)

func main() {
	if err := cmd.NewRootCommand().Execute(); err != nil {
		os.Exit(-1)
	}
}
