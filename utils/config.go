package utils

import (
	"flag"

	"github.com/robfig/config"
)

var (
	MYSQL_SECTION = "MYSQL"
	CONF_FILE     = flag.String("confFile", "conf/app.conf", "General configuration file")
	//MYSQL_DSN_DATA = "root:923923924@tcp(127.0.0.1:3306)/mydb?charset=utf8&parseTime=true&loc=Local"
)

type Setting struct {
	Driver  string
	DSNUser string
	DSNData string
}

func GetSetting() *Setting {
	c, _ := config.ReadDefault(*CONF_FILE)

	driver, _ := c.String(config.DEFAULT_SECTION, "mysql_driver")
	dsnUser, _ := c.String(MYSQL_SECTION, "dsn_user")
	dsnData, _ := c.String(MYSQL_SECTION, "dsn_data")

	return &Setting{driver, dsnUser, dsnData}
}
