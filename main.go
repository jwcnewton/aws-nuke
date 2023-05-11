package main

import (
	"os"

	"github.com/jwcnewton/aws-nuke/v3/cmd"
)

func main() {
	if err := cmd.NewRootCommand().Execute(); err != nil {
		os.Exit(-1)
	}
}
