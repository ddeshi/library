package api

import (
	"github.com/ddeshi/library/pkg/Service/User"
	"github.com/gin-gonic/gin"
)

type UserAPIController struct{}

func UserRegisterRouters(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.POST("/login", User.Login)
		user.POST("/register", User.Register)
	}
}
