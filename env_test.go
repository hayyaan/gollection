package gollection

import (
	"github.com/sosedoff/pgweb/Godeps/_workspace/src/github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	os.Clearenv()
	assert.Equal(t, "bar", GetEnv("foo", "bar"))

	os.Setenv("foo", "baz")
	assert.Equal(t, "baz", GetEnv("foo", "bar"))
}

func TestGetEnvInt(t *testing.T) {
	os.Clearenv()
	assert.Equal(t, 42, GetEnvInt("foo", 42))

	os.Setenv("foo", "1234")
	assert.Equal(t, 1234, GetEnvInt("foo", 42))
}
