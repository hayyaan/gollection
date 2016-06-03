package database

import (
	"fmt"

	"github.com/gollection/gollection"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// New returns a new mysql gorm db connection
// Set connection config in the gollection config
func New(c gollection.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@%s:%d/%s?charset=utf8&parseTime=True&loc=Local",
		c.DBConfig.Username,
		c.DBConfig.Password,
		c.DBConfig.Host,
		c.DBConfig.Port,
		c.DBConfig.Database,
	)

	return gorm.Open("mysql", dsn)
}
