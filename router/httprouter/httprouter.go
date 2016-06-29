package httprouter

import (
	"fmt"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/gollection/gollection/router"
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/cli"
)

type httprouterWrapper struct {
	config router.Config
	Router *httprouter.Router
}

func New(logger log.Logger, c router.Config) *httprouterWrapper {
	return &httprouterWrapper{
		config: c,
		Router: httprouter.New(),
	}
}

func (r httprouterWrapper) Cli() cli.Command {
	addr := fmt.Sprintf("%s:%d", r.config.Host, r.config.Port)

	return cli.Command{
		Name:  "serve",
		Usage: "Run the http server that listens on " + addr,
		Action: func(c *cli.Context) error {
			return http.ListenAndServe(addr, r.Router)
		},
	}
}
