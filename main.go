package main

import (
	"github.com/spf13/pflag"
	//gin "github.com/gin-gonic/gin"
	"github.com/lfuture/easygin/pkg/app"
	"fmt"
	"github.com/lfuture/easygin/pkg/database/mysql"
	"github.com/lfuture/easygin/pkg/database/redis"
)

type User struct {
	ID int `gorm:"column:id"`
	Mapkey string
	Mapvalue string
}

func (User) TableName() string {
	return "my_password"
}

func main() {
	//gin.Default()
	var env *string = pflag.String("env","", "运行环境类型：sandbox or production")
	pflag.Parse()
	if *env != "sandbox" && *env != "production" && *env != "development" {
		pflag.PrintDefaults()
		return
	}

	app.Bootstrap(*env)

	// config test
	fmt.Println(app.GetConfig().Section("").Key("name"))


	mysql.Bootstrap()

	// mysql test
	var user = new(User)
	mysql.DB.First(user, 6)

	fmt.Println(user.ID)
	fmt.Println(user.Mapkey)
	fmt.Println(user.Mapvalue)

	redis.Bootstrap()

	result := redis.RedisClient.Get("liu")
	println(result.Val())
}
