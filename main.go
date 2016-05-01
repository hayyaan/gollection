package gollection

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"
)

// Gollection holds everything for your application to work
type Gollection struct {
	Cli    *cli.App
	Env    Env
	Router *gin.Engine
}

// New creates a new gollection with minimum requirements
func New() *Gollection {
	gollection := Gollection{
		Cli: cli.NewApp(),
	}

	return &gollection
}

// Run runs your gollection application. Fingers crossed
func (gollection *Gollection) Run() error {
	return gollection.Cli.Run(os.Args)
}

// SetEnv gets your applications Env and passes them to gollection
func (gollection *Gollection) SetEnv(env Env) {
	gollection.Env = env
}
