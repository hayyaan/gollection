package gollection

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseEnvFile(t *testing.T) {
	assert.Equal(t,
		map[string]string{},
		parseEnvFile("\n\n"),
	)

	assert.Equal(t,
		map[string]string{"APP_HOST": "0.0.0.0", "APP_PORT": "1234"},
		parseEnvFile("APP_HOST=0.0.0.0\nAPP_PORT=1234\n\n"),
	)

	assert.Equal(t,
		map[string]string{},
		parseEnvFile("foo\none=two=tree\n\n"),
	)
}

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

func TestGetEnvBool(t *testing.T) {
	os.Clearenv()
	assert.True(t, GetEnvBool("foo", true))
	assert.False(t, GetEnvBool("foo", false))

	os.Setenv("foo", "true")
	assert.True(t, GetEnvBool("foo", true))

	os.Setenv("foo", "false")
	assert.False(t, GetEnvBool("foo", true))
}
