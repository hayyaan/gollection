package sqlite

import (
	"github.com/gollection/gollection"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func New(c gollection.Config) (*gorm.DB, error) {
	return gorm.Open("sqlite3", c.DBConfig.Database)
}
