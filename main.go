package gollection

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

type Gollection struct {
	Config  interface{}
	Router  *gin.Engine
	Command cobra.Command
}

func New() *Gollection {
	return &Gollection{}
}

func (gollection *Gollection) Run() error {
	return gollection.Router.Run(":1234")
}
