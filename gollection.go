package gollection

import (
	"os"

	"github.com/gollection/gollection/cache"
	"github.com/gollection/gollection/log"
	"github.com/gollection/gollection/router"
	"github.com/jinzhu/gorm"
	"github.com/urfave/cli"
	"gopkg.in/redis.v3"
)

// Gollection holds everything for your application to work
type Gollection struct {
	Cache        cache.Cache
	Cli          *cli.App
	Config       Config
	DB           *gorm.DB
	Log          log.Logger
	Redis        *redis.Client
	RouterEngine router.Engine
}

// New creates a new gollection with minimum requirements
func New(c Config) *Gollection {
	g := Gollection{
		Cli:    cli.NewApp(),
		Config: c,
		Log:    log.NewLog15(),
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
		g.Log.Crit("Can't connect to the DB", "err", err)
	}

	if g.Config.AppConfig.Debug {
		db.LogMode(true)
		db.SetLogger(log.NewGormLogger(g.Log))
	}

	g.DB = db
}

func (g *Gollection) AddRedis(r *redis.Client, err error) {
	if err != nil {
		g.Log.Crit("Can't connect to Redis", err)
	}

	err = r.Ping().Err()
	if err != nil {
		g.Log.Crit("Can't connect to Redis", err)
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
