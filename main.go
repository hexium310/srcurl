package main

import (
	"os"

	"github.com/hexium310/srcurl/cmd"
)

func main() {
	err := command.RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
