package gin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGin_New(t *testing.T) {
	r := New()
	assert.NotNil(t, r)
	assert.NotNil(t, r.Engine)
	//TODO: Test something real for this?!
}

func TestGin_Cli(t *testing.T) {
	r := New()
	cli := r.Cli()
	assert.Equal(t, "serve", cli.Name)
	assert.Equal(t, "Run the http server that listens on ", cli.Usage)
	assert.NotNil(t, cli.Action) // TODO: Test more than if Action != nil
}