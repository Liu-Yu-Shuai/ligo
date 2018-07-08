package logging

import (
	"github.com/lestrrat-go/file-rotatelogs"
	config2 "github.com/yushuailiu/easygin/pkg/config"
)

func GetRequestLogger() *rotatelogs.RotateLogs {
	config := config2.GetConfig().Section("logging")

	logName := config.Key("requestLog").String()

	return rotateWriter(logName)
}
