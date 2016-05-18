package gollection

import (
	"log"
	"os"

	"github.com/MetalMatze/gollection/cache"
	"github.com/MetalMatze/gollection/router"
	"github.com/codegangsta/cli"
	"github.com/jinzhu/gorm"
	"gopkg.in/redis.v3"
)

// Gollection holds everything for your application to work
type Gollection struct {
	Cache        cache.Cache
	Cli          *cli.App
	Config       Config
	DB           *gorm.DB
	Redis        *redis.Client
	RouterEngine router.Engine
}

// New creates a new gollection with minimum requirements
func New(c Config) *Gollection {
	g := Gollection{
		Cli:    cli.NewApp(),
		Config: c,
	}

	g.startCli()

	return &g
}

// Run runs your gollection application. Fingers crossed
func (g *Gollection) Run() error {
	return g.Cli.Run(os.Args)
}

func (g *Gollection) AddDB(db *gorm.DB, err error) {
	if err != nil {
		log.Fatal(err)
	}

	if g.Config.AppConfig.Debug {
		db.LogMode(true)
	}

	g.DB = db
}

func (g *Gollection) AddRedis(r *redis.Client, err error) {
	if err != nil {
		log.Fatal(err)
	}

	err = r.Ping().Err()
	if err != nil {
		log.Fatal(err)
	}

	g.Redis = r
}

func (g *Gollection) AddCache(c cache.Cache) {
	g.Cache = c
}

func (g *Gollection) AddRouter(e router.Engine) {
	g.RouterEngine = e
	g.RouterEngine.Debug(g.Config.AppConfig.Debug)
}

func (g *Gollection) AddRoutes(r func(router.Router)) {
	r(g.RouterEngine.Router())
}
