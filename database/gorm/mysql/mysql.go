package mysql

import (
	"fmt"

	"github.com/go-kit/kit/log"
	gogorm "github.com/gollection/gollection/database/gorm"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// New returns a new mysql gorm db connection
func New(logger log.Logger, c gogorm.Config) (*gorm.DB, error) {
	g, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@%s:%d/%s?charset=utf8&parseTime=True&loc=Local",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
	))

	g.SetLogger(gogorm.NewLogger(logger))

	return g, err
}
