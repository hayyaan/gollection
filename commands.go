package gollection

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/rubenv/sql-migrate"
	"github.com/urfave/cli"
)

func (g *Gollection) startCli() {
	g.Cli.Name = g.Config.AppConfig.Name
	g.Cli.Usage = g.Config.AppConfig.Usage

	g.addDBCommand()
}

// AddCommands takes commands and adds them additionally to gollection's commands.
func (gollection *Gollection) AddCommands(commands ...cli.Command) {
	for _, command := range commands {
		gollection.Cli.Commands = append(gollection.Cli.Commands, command)
	}
}

func (g *Gollection) addServeCommand() {
	addr := fmt.Sprintf("%s:%d", g.Config.AppConfig.Host, g.Config.AppConfig.Port)
	g.Cli.Commands = append(g.Cli.Commands, cli.Command{
		Name:  "serve",
		Usage: "Run the http server that listens on " + addr,
		Action: func(c *cli.Context) error {
			if g.Config.Debug {
				go func() {
					http.ListenAndServe("localhost:6060", nil)
				}()
			}

			return g.RouterEngine.Run(addr)
		},
	})
}

func (g *Gollection) addDBCommand() {
	migrations := &migrate.FileMigrationSource{Dir: "database/migrations"}

	g.Cli.Commands = append(g.Cli.Commands, cli.Command{
		Name:  "migrate",
		Usage: "Run migration actions",
		Subcommands: []cli.Command{{
			Name:  "up",
			Usage: "Migrate your database",
			Action: func(c *cli.Context) error {
				return migrateUp(g, migrations)
			},
		}, {
			Name:  "down",
			Usage: "Rollback all database migrations",
			Action: func(c *cli.Context) error {
				return migrateDown(g, migrations)
			},
		}, {
			Name:  "reset",
			Usage: "Reset the database by running down & up migrations",
			Action: func(c *cli.Context) error {
				if err := migrateDown(g, migrations); err != nil {
					return err
				}

				if err := migrateUp(g, migrations); err != nil {
					return err
				}

				return nil
			},
		}},
	})
}

func migrateUp(g *Gollection, m *migrate.FileMigrationSource) error {
	applied, err := migrate.Exec(g.DB.DB(), g.Config.DBConfig.Dialect, m, migrate.Up)
	if err != nil {
		g.Log.Warn("Can't run migrations", "err", err)
	}
	g.Log.Info(fmt.Sprintf("Applied to %d migrations!", applied))

	return nil
}

func migrateDown(g *Gollection, m *migrate.FileMigrationSource) error {
	applied, err := migrate.Exec(g.DB.DB(), g.Config.DBConfig.Dialect, m, migrate.Down)
	if err != nil {
		g.Log.Warn("Can't run migrations", "err", err)
	}

	g.Log.Info(fmt.Sprintf("Rolled back %d migrations!", applied))

	return nil
}
