package handlers

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/lfuture/easygin/pkg/logging"
	"github.com/lfuture/easygin/pkg/http"
)

type User struct {
	FirstName string `form:"first_name" json:"first_name" binding:"required,email"`
	LastName string `form:"last_name" json:"last_name" binding:"required"`
}

func AddUserHandler(c *gin.Context)  {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		logging.Log.Error(err)
		http.Fail(c, "", nil)
		return
	}

	fmt.Println(user)
}