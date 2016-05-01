package gollection

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"
)

type Gollection struct {
	Cli    *cli.App
	Env    Env
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

func (gollection *Gollection) SetEnv(env GollectionEnv) {
	gollection.Env = env
}
