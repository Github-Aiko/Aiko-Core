package main

import (
	"fmt"

	"github.com/Github-Aiko/Aiko-Core/core"
	"github.com/Github-Aiko/Aiko-Core/main/commands/base"
)

var cmdVersion = &base.Command{
	UsageLine: "{{.Exec}} version",
	Short:     "Show current version of Aiko",
	Long: `Version prints the build information for Aiko executables.
	`,
	Run: executeVersion,
}

func executeVersion(cmd *base.Command, args []string) {
	printVersion()
}

func printVersion() {
	version := core.VersionStatement()
	for _, s := range version {
		fmt.Println(s)
	}
}
