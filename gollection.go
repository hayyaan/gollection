package gollection

import (
	"github.com/urfave/cli"
	"os"
)

// Gollection is the base object for everything
type Gollection struct {
	cli    *cli.App
	config Config
}

// New creates a new Gollection object from the given config and instantiates the cli app
func New(c Config) *Gollection {
	g := &Gollection{
		cli:    cli.NewApp(),
		config: c,
	}

	g.cli.Name = g.config.Name
	g.cli.Usage = g.config.Usage

	return g
}

func (g *Gollection) Run() error {
	return g.cli.Run(os.Args)
}

// Register takes different kinds of providers and registers them correctly with gollection
func (g *Gollection) Register(s interface{}) {
	cliService, ok := s.(CLIProvider)
	if ok {
		g.cli.Commands = append(g.cli.Commands, cliService.Cli())
	}
}
