package postgres

import (
	"github.com/go-kit/kit/log"
	"github.com/gollection/gollection/database"
	gogorm "github.com/gollection/gollection/database/gorm"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// New returns a new postgres gorm db connection
func New(logger log.Logger, c database.Config) (*gorm.DB, error) {
	g, err := gorm.Open("postgres", c.DSN()+"?sslmode=disable")

	g.SetLogger(gogorm.NewLogger(logger))

	return g, err
}
