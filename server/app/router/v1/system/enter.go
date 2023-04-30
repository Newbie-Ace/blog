package system

import (
	"github.com/gin-gonic/gin"
	"github.com/noobHuKai/app/middleware"
)

type RouterGroup struct {
	UserRouter
	CommonRouter
}

func (r RouterGroup) InitSystemRouter(router *gin.RouterGroup) {
	commonRouter := router.Group("")
	{
		r.InitCommonRouter(commonRouter)
	}

	// after interface must with token
	router.Use(middleware.TokenAuthorizeMiddleware())

	userRouter := router.Group("user")
	{
		r.InitUserRouter(userRouter)
	}
}
