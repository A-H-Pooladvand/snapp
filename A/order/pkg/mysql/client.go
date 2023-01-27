package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"order/configs"
)

func New() *gorm.DB {
	c := &client{
		configs.NewMysql(),
	}

	return c.connect()
}

type client struct {
	configs.Mysql
}

func (c *client) connect() (db *gorm.DB) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       c.dsn(), // data source name
		DefaultStringSize:         256,     // default size for string fields
		DisableDatetimePrecision:  true,    // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,    // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,    // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,   // autoconfigure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return
}

func (c *client) dsn() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.User,
		c.Pass,
		c.Host,
		c.Port,
		c.DATABASE,
	)
}
