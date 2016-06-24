package gollection

import "github.com/urfave/cli"

// CLIProvider is an interface that adds its cli.Command to gollection's root cli command
type CLIProvider interface {
	Cli() cli.Command
}
