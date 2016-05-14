package gollection

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
	"github.com/rubenv/sql-migrate"
)

func (g *Gollection) startCli() {
	g.Cli.Name = g.Config.AppConfig.Name
	g.Cli.Usage = g.Config.AppConfig.Usage
	g.Cli.EnableBashCompletion = true

	g.addServeCommand()
	g.addDBCommand()
}

// AddCommands takes commands and adds them additionally to gollection's commands.
func (gollection *Gollection) AddCommands(commands ...cli.Command) {
	for _, command := range commands {
		gollection.Cli.Commands = append(gollection.Cli.Commands, command)
	}
}

func (gollection *Gollection) addServeCommand() {
	addr := fmt.Sprintf("%s:%d", gollection.Config.AppConfig.Host, gollection.Config.AppConfig.Port)
	gollection.Cli.Commands = append(gollection.Cli.Commands, cli.Command{
		Name:  "serve",
		Usage: "Run the http server that listens on " + addr,
		Action: func(c *cli.Context) {
			gollection.RouterEngine.Run(addr) // TODO: Return the error
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
			Action: func(c *cli.Context) {
				applied, err := migrate.Exec(g.DB.DB(), g.Config.DBConfig.Dialect, migrations, migrate.Up)
				if err != nil {
					log.Fatal(err)
				}
				log.Printf("Applied to %d migrations!\n", applied)
			},
		}, {
			Name:  "down",
			Usage: "Rollback all database migrations",
			Action: func(c *cli.Context) {
				applied, err := migrate.Exec(g.DB.DB(), g.Config.DBConfig.Dialect, migrations, migrate.Down)
				if err != nil {
					log.Fatal(err)
				}
				log.Printf("Rolled back %d migrations!\n", applied)
			},
		}},
	})
}
