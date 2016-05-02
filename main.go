package gollection

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"
)

// Gollection holds everything for your application to work
type Gollection struct {
	Cli    *cli.App
	Config Config
	Router *gin.Engine
}

// New creates a new gollection with minimum requirements
func New(config Config) *Gollection {
	gin.SetMode(gin.ReleaseMode)

	gollection := Gollection{
		Cli:    cli.NewApp(),
		Config: config,
		Router: gin.Default(),
	}

	gollection.startCli()

	return &gollection
}

// Run runs your gollection application. Fingers crossed
func (gollection *Gollection) Run() error {
	return gollection.Cli.Run(os.Args)
}

func (gollection *Gollection) AddRoutes(routes func(*gin.Engine)) {
	routes(gollection.Router)
}
