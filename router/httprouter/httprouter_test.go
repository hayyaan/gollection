package httprouter

import (
	"bytes"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/gollection/gollection/router"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	logger := log.NewLogfmtLogger(&buf)

	r := New(logger, router.Config{})
	assert.NotNil(t, r)
	assert.NotNil(t, r.Router)
}

func TestHttprouterWrapper_Cli(t *testing.T) {
	var buf bytes.Buffer
	logger := log.NewLogfmtLogger(&buf)

	r := New(logger, router.Config{Host: "foobar", Port: 666})

	cli := r.Cli()
	assert.Equal(t, "serve", cli.Name)
	assert.Equal(t, "Run the http server that listens on foobar:666", cli.Usage)
}
