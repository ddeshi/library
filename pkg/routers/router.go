package routers

import (
	"github.com/ddeshi/library/pkg/routers/api"
	"github.com/gin-gonic/gin"
)

type APIController interface {
	RegisterRouters(r *gin.RouterGroup)
}

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(
		gin.Logger(),
		gin.Recovery(),
	)
	// 相关API操作
	api.UserRegisterRouters(r)
	api.BookRegisterRouters(r)

	return r
}
