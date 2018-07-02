package mysql

import (
	"github.com/lfuture/easygin/pkg/app"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"github.com/go-ini/ini"
)


var DB *gorm.DB

func Bootstrap()  {
	mysqlConfig, err := app.GetConfig().GetSection("mysql")
	if err != nil {
		panic(err)
	}

	initMysql(mysqlConfig)
}

func initMysql(config *ini.Section)  {

	url := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", config.Key("username").String(),
		config.Key("password").String(),
		config.Key("host").String(), config.Key("dbname"))
	gormDB, err := gorm.Open("mysql", url)

	maxIdle,err := config.Key("maxIdle").Int()

	if err != nil {
		panic(err)
	}

	gormDB.DB().SetMaxIdleConns(maxIdle)
	if err != nil {
		panic(err)
	}
	DB =gormDB
}