package gollection

import (
	"os"

	"github.com/go-kit/kit/log"
	"github.com/urfave/cli"
)

// Gollection is the base object for everything
type Gollection struct {
	Cli    *cli.App
	Config Config
	Logger log.Logger
}

// New creates a new Gollection object from the given config and instantiates the cli app
func New(c Config) *Gollection {
	g := &Gollection{
		Cli:    cli.NewApp(),
		Config: c,
	}

	g.Cli.Name = g.Config.Name
	g.Cli.Usage = g.Config.Usage

	return g
}

func (g *Gollection) Run() error {
	return g.Cli.Run(os.Args)
}

// Register takes different kinds of providers and registers them correctly with gollection
func (g *Gollection) Register(s interface{}) {
	cliService, ok := s.(CLIProvider)
	if ok {
		g.Cli.Commands = append(g.Cli.Commands, cliService.Cli())
	}
}
