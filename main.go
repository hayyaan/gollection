package gollection

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"
)

type Gollection struct {
	Cli    *cli.App
	Config GollectionConfig
	Router *gin.Engine
}

func New() *Gollection {
	gollection := Gollection{
		Cli: cli.NewApp(),
	}

	return &gollection
}

func (gollection *Gollection) Run() error {
	return gollection.Cli.Run(os.Args)
}

func (gollection *Gollection) SetConfig(gc GollectionConfig) {
	gollection.Config = gc
}
