package main

import (
	"os"

	"gitlab.com/MarkCherepovskyi/KeyStorage/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
