package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_DSN(t *testing.T) {
	c := Config{Driver: "postgres", Username: "user", Password: "pass", Host: "host", Port: 5432, Database: "database"}
	assert.Equal(t, `postgres://user:pass@host:5432/database`, c.DSN())

	c = Config{Driver: "mysql", Username: "user", Password: "pass", Host: "host", Port: 3306, Database: "database"}
	assert.Equal(t, `mysql://user:pass@host:3306/database`, c.DSN())

	c = Config{Driver: "sqlite", Database: "/opt/databases/mydb.sq3"}
	assert.Equal(t, `sqlite:/opt/databases/mydb.sq3`, c.DSN())
}
