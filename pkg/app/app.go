package app

import (
	"github.com/go-ini/ini"
	"github.com/lfuture/easygin/pkg/config"
	"github.com/lfuture/easygin/routes"
	"github.com/gin-gonic/gin"
)

type App struct {
	env string
	mode string
	Config *ini.File
}

var app = &App{}

func Bootstrap(env string) *gin.Engine {
	app.env = env

	app.Config = config.DefaultConfig().Bootstrap(env)

	server := routes.InitRoute()
	return server
}

func GetConfig() *ini.File {
	return app.Config
}

func Env() string {
	return app.env
}

func IsDebug() bool {
	return app.mode == "debug"
}
