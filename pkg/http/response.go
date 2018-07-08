package http

import (
	"github.com/gin-gonic/gin"
	"github.com/yushuailiu/easygin/pkg/err"
	"net/http"
)

func Fail(c *gin.Context, msg string, data interface{}) {

	if msg == "" {
		msg = err.GetCodeMsg(http.StatusBadRequest)
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"msg":	msg,
		"data":	data,
	})
}

func Success(c *gin.Context, msg string, data interface{}) {

	if msg == "" {
		msg = err.GetCodeMsg(http.StatusOK)
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":	msg,
		"data":	data,
	})
}

func SystemError(c *gin.Context, msg string, data interface{}) {

	if msg == "" {
		msg = err.GetCodeMsg(http.StatusInternalServerError)
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"msg":	msg,
		"data":	data,
	})
}
