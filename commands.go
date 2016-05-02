package gollection

import (
	"fmt"

	"github.com/codegangsta/cli"
)

func (gollection *Gollection) startCli() {
	gollection.Cli.Name = gollection.Config.AppConfig.Name
	gollection.Cli.Usage = gollection.Config.AppConfig.Usage
	gollection.Cli.EnableBashCompletion = true

	gollection.addServeCommand()
}

// AddCommands takes commands and adds them additionally to gollection's commands.
func (gollection *Gollection) AddCommands(commands ...cli.Command) {
	for _, command := range commands {
		gollection.Cli.Commands = append(gollection.Cli.Commands, command)
	}
}

func (gollection *Gollection) addServeCommand() {
	addr := fmt.Sprintf("%s:%d", gollection.Config.Host, gollection.Config.Port)
	gollection.Cli.Commands = append(gollection.Cli.Commands, cli.Command{
		Name:  "serve",
		Usage: "Run the http server that listens on " + addr,
		Action: func(c *cli.Context) {
			gollection.Router.Run(addr) // TODO: Return the error
		},
	})
}
