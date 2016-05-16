package database

import (
	"github.com/MetalMatze/gollection"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func SQLite(c gollection.Config) (*gorm.DB, error) {
	return gorm.Open("sqlite3", c.DBConfig.Database)
}
