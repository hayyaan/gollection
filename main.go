package gollection

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"
)

// Gollection holds everything for your application to work
type Gollection struct {
	Cli    *cli.App
	Router *gin.Engine
}

// New creates a new gollection with minimum requirements
func New() *Gollection {
	gin.SetMode(gin.ReleaseMode)

	gollection := Gollection{
		Cli:    cli.NewApp(),
		Router: gin.New(),
	}

	return &gollection
}

// Run runs your gollection application. Fingers crossed
func (gollection *Gollection) Run() error {
	return gollection.Cli.Run(os.Args)
}

func (gollection *Gollection) AddRoutes(routes func(*gin.Engine)) {
	routes(gollection.Router)
}
