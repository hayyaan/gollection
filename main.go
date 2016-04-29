package gollection

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

type Gollection struct {
	Config  GollectionConfig
	Router  *gin.Engine
	Command cobra.Command
}

func New() *Gollection {
	return &Gollection{}
}

func (gollection *Gollection) SetConfig(gc GollectionConfig) {
	gollection.Config = gc
}

func (gollection *Gollection) Run() error {
	addr, port := gollection.Config.GetAddrPort()
	return gollection.Router.Run(fmt.Sprintf("%s:%d", addr, port))
}
