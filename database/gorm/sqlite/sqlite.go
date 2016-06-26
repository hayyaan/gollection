package sqlite

import (
	"github.com/go-kit/kit/log"
	gogorm "github.com/gollection/gollection/database/gorm"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// New returns a new sqlite gorm db connection
func New(logger log.Logger, c gogorm.Config) (*gorm.DB, error) {
	g, err := gorm.Open("sqlite3", c.Database)
	g.SetLogger(gogorm.NewLogger(logger))

	return g, err
}
