package database

import "fmt"

const DriverPostgres string = "postgres"
const DriverMysql string = "mysql"
const DriverSqlite string = "sqlite"

// Config describes how to connect to a database
type Config struct {
	Driver   string
	Username string
	Password string
	Host     string
	Port     int
	Database string
}

func (c Config) DSN() string {
	switch c.Driver {
	case DriverPostgres, DriverMysql:
		return fmt.Sprintf("%s://%s:%s@%s:%d/%s",
			c.Driver,
			c.Username,
			c.Password,
			c.Host,
			c.Port,
			c.Database,
		)
	case DriverSqlite:
		return "sqlite:" + c.Database
	default:
		return ""
	}
}
