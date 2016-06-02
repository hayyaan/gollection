package postgres

import (
	"fmt"

	"github.com/gollection/gollection"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func New(c gollection.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		c.DBConfig.Username,
		c.DBConfig.Password,
		c.DBConfig.Host,
		c.DBConfig.Port,
		c.DBConfig.Database,
	)

	return gorm.Open("postgres", dsn)
}
