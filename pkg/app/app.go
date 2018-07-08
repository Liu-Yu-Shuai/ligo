package app

import (
	"github.com/go-ini/ini"
	"github.com/yushuailiu/easygin/pkg/config"
	"github.com/yushuailiu/easygin/pkg/logging"
	"github.com/yushuailiu/easygin/routes"
	"github.com/gin-gonic/gin"
	"io"
	"os"
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

	if IsDevelopment() {
		gin.DefaultWriter = io.MultiWriter(logging.GetRequestLogger(), os.Stdout)
	} else {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = logging.GetRequestLogger()
	}

	server := routes.InitRoute()
	return server
}

func GetConfig() *ini.File {
	return app.Config
}

func Env() string {
	return app.env
}

func IsDevelopment() bool {
	return app.env == "development"
}

func IsDebug() bool {
	return app.mode == "debug"
}
