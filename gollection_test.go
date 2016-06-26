package gollection

import (
	"bytes"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
)

var c = Config{Name: "app", Usage: "my app"}

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	logger := log.NewLogfmtLogger(&buf)

	g := New(logger, c)

	assert.Equal(t, c, g.Config)
	assert.NotNil(t, g.Cli)
	assert.IsType(t, &cli.App{}, g.Cli)
	assert.Equal(t, c.Name, g.Cli.Name)
	assert.Equal(t, c.Usage, g.Cli.Usage)
}
