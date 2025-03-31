package main

import (
	"github.com/the-mold/go-project-layout/cmd"

	// must be added to generate proper cli tool
	_ "github.com/the-mold/go-project-layout/cmd/server"
)

func main() {
	cmd.Execute()
}
