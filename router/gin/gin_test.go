package gin

import (
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/gollection/gollection/router"
	"github.com/stretchr/testify/assert"
)

func TestGin_New(t *testing.T) {
	r := New(log.NewNopLogger(), router.Config{})
	assert.NotNil(t, r)
	assert.NotNil(t, r.Engine)
	//TODO: Test something real for this?!
}

func TestGin_Cli(t *testing.T) {
	r := New(log.NewNopLogger(), router.Config{Host: "foobar", Port: 666})
	cli := r.Cli()
	assert.Equal(t, "serve", cli.Name)
	assert.Equal(t, "Run the http server that listens on foobar:666", cli.Usage)
	assert.NotNil(t, cli.Action) // TODO: Test more than if Action != nil
}
