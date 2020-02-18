package utils

import (
	"flag"
	"fmt"
	"sync"

	"github.com/robfig/config"
)

var (
	configFile   = flag.String("conf_file", "conf/app.conf", "General configuration file")
	mysqlSection = "MYSQL"

	mu         sync.Mutex // protects settings
	appSetting *Setting
)

// Setting 应用配置
type Setting struct {
	Driver  string
	DSNUser string
	DSNData string
}

func getSetting() *Setting {
	mu.Lock()
	defer mu.Unlock()

	if appSetting == nil {
		c, _ := config.ReadDefault(*configFile)
		driver, _ := c.String(config.DEFAULT_SECTION, "mysql_driver")
		dsnUser, _ := c.String(mysqlSection, "dsn_user")
		dsnData, _ := c.String(mysqlSection, "dsn_data")
		appSetting = &Setting{driver, dsnUser, dsnData}
		fmt.Println("app setting：", appSetting)
	}
	return appSetting
}

// GetSetting 获取应用配置
func GetSetting() *Setting {
	if appSetting != nil {
		return appSetting
	}
	return getSetting()
}

func init() {
	flag.Parse()
}
