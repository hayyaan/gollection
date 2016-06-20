package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
)

// Gin is a gin-gonic wrapper that implements gollection's interfaces
type Gin struct {
	Engine *gin.Engine
}

// New creates a new Gin wrapper
func New() *Gin {
	return &Gin{
		Engine: gin.Default(),
	}
}

// Cli implements the CliProvider interface and adds gin to the set of cli commands
func (g Gin) Cli() cli.Command {
	return cli.Command{
		Name:  "serve",
		Usage: "Run the http server that listens on ", //+ addr,
		Action: func(c *cli.Context) error {
			g.Engine.Run(":1234")
			return nil
		},
	}
}
