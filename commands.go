package gollection

import (
	"fmt"

	"github.com/codegangsta/cli"
)

func (gollection *Gollection) AddCommands() {
	host, port := gollection.Config.GetHostPort()
	addr := fmt.Sprintf("%s:%d", host, port)
	gollection.Cli.Commands = append(gollection.Cli.Commands, cli.Command{
		Name:  "serve",
		Usage: "Run the http server that listens on " + addr,
		Action: func(c *cli.Context) {
			gollection.Router.Run(addr) // TODO: Return the error
		},
	})
}
