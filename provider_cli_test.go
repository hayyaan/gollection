package gollection

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
)

var command = cli.Command{
	Name:  "msg",
	Usage: "Prints a message",
	Action: func(c *cli.Context) error {
		return nil
	},
}

type MessageService struct {
	Msg string
}

func (ms MessageService) Cli() cli.Command {
	return command
}

func TestCli(t *testing.T) {
	g := New(Config{Name: "app", Usage: "Run an app"})
	assert.NotNil(t, g)
	assert.NotNil(t, g.Cli)
	assert.Equal(t, "app", g.Cli.Name)
	assert.Equal(t, "Run an app", g.Cli.Usage)
	assert.Len(t, g.Cli.Commands, 0)

	g.Register(MessageService{Msg: "foobar"})
	assert.Len(t, g.Cli.Commands, 1)
	assert.ObjectsAreEqual(command, g.Cli.Commands[0])
}
