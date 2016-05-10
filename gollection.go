package gollection

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Gollection holds everything for your application to work
type Gollection struct {
	Cli    *cli.App
	Config Config
	DB     *gorm.DB
	Router *gin.Engine
}

// New creates a new gollection with minimum requirements
func New(c Config) *Gollection {
	gin.SetMode(gin.ReleaseMode)

	g := Gollection{
		Cli:    cli.NewApp(),
		Config: c,
		Router: gin.Default(),
	}

	g.startCli()

	return &g
}

// Run runs your gollection application. Fingers crossed
func (g *Gollection) Run() error {
	return g.Cli.Run(os.Args)
}

func (g *Gollection) AddRoutes(routes func(*gin.Engine)) {
	routes(g.Router)
}

func (g *Gollection) AddDB(db *gorm.DB, err error) {
	if err != nil {
		log.Fatal(err)
	}

	if g.Config.AppConfig.Env == "local" {
		db.LogMode(true)
	}

	g.DB = db
}
