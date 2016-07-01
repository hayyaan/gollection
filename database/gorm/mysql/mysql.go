package mysql

import (
	"github.com/go-kit/kit/log"
	"github.com/gollection/gollection/database"
	gogorm "github.com/gollection/gollection/database/gorm"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// New returns a new mysql gorm db connection
func New(logger log.Logger, c database.Config) (*gorm.DB, error) {
	g, err := gorm.Open("mysql", c.DSN()+"?charset=utf8&parseTime=True&loc=Local")

	g.SetLogger(gogorm.NewLogger(logger))

	return g, err
}
