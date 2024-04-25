package database

import (
	"brew-note/src/config"

	mysqldriver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

const net = "tcp"

func init() {
	c := mysqldriver.NewConfig()
	c.DBName = config.DB.Name
	c.User = config.DB.User
	c.Passwd = config.DB.Password
	c.Addr = config.DB.Host
	c.Net = net
	c.ParseTime = true

	db, err := gorm.Open(mysql.Open(c.FormatDSN()))
	if err != nil {
		panic(err.Error())
	}

	DB = db
}