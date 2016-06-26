package gorm

import (
	"testing"

	"bytes"

	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/assert"
)

func TestLogger_Print(t *testing.T) {
	var buf bytes.Buffer
	logger := log.NewLogfmtLogger(&buf)
	l := NewLogger(logger)

	assert.Equal(t, logger, l.logger)

	l.Print("foo", "path", "duration", "query", "param")
	assert.Equal(t, "path=path duration=duration query=query param=param\n", buf.String())
}
