package gin

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/log"
	"github.com/urfave/cli"
)

type GinConfig struct {
	Host string
	Port int
}

// Gin is a gin-gonic wrapper that implements gollection's interfaces
type Gin struct {
	config GinConfig
	Engine *gin.Engine
}

// New creates a new Gin wrapper
func New(logger log.Logger, c GinConfig) *Gin {
	gin.SetMode(gin.ReleaseMode)

	g := &Gin{
		config: c,
		Engine: gin.New(),
	}

	g.Engine.Use(gin.Recovery())
	g.Engine.Use(ginLogging(logger))

	return g
}

// Cli implements the CliProvider interface and adds gin to the set of cli commands
func (g Gin) Cli() cli.Command {
	return cli.Command{
		Name:  "serve",
		Usage: "Run the http server that listens on ", //+ addr,
		Action: func(c *cli.Context) error {
			g.Engine.Run(fmt.Sprintf("%s:%d", g.config.Host, g.config.Port))
			return nil
		},
	}
}

// This is the official gin logger yet a bit changed to use go-kit's log.Logger interface
func ginLogging(logger log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path

		// Process request
		c.Next()

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()

		logger.Log(
			"status", statusCode,
			"duration", time.Since(start),
			"method", method,
			"ip", clientIP,
			"path", path,
		)
	}
}
