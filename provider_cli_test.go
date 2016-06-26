package gollection

import (
	"bytes"
	"testing"

	"github.com/go-kit/kit/log"
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
	var buf bytes.Buffer
	logger := log.NewLogfmtLogger(&buf)

	g := New(logger, Config{Name: "app", Usage: "Run an app"})
	assert.NotNil(t, g)
	assert.NotNil(t, g.Cli)
	assert.Equal(t, "app", g.Cli.Name)
	assert.Equal(t, "Run an app", g.Cli.Usage)
	assert.Len(t, g.Cli.Commands, 0)

	g.Register(MessageService{Msg: "foobar"})
	assert.Len(t, g.Cli.Commands, 1)
	assert.ObjectsAreEqual(command, g.Cli.Commands[0])
}
