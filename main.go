package main

import (
	"github.com/spf13/pflag"
	//gin "github.com/gin-gonic/gin"
	"github.com/lfuture/easygin/pkg/app"
	"fmt"
	"github.com/lfuture/easygin/pkg/database/mysql"
	"github.com/lfuture/easygin/pkg/database/redis"
	"github.com/lfuture/easygin/app/models"
	"github.com/lfuture/easygin/pkg/logging"
)


func main() {
	//gin.Default()
	var env *string = pflag.String("env","", "运行环境类型：development、sandbox、production")
	pflag.Parse()
	if *env != "sandbox" && *env != "production" && *env != "development" {
		pflag.PrintDefaults()
		return
	}

	server := app.Bootstrap(*env)

	// config test
	fmt.Println(app.GetConfig().Section("").Key("name"))


	mysql.Bootstrap()


	redis.Bootstrap()

	result := redis.RedisClient.Get("liu")
	println(result.Val())

	models.AddUser("shuai", "liu")

	server.BasePath()

	logging.Bootstrap()

	logging.Log.Error("liushuai~")
	logging.GetLogger("test").Info("asdfasdfasdfsadfasdf")
}
