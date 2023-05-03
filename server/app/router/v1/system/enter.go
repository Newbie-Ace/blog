package system

import (
	"github.com/gin-gonic/gin"
	"github.com/noobHuKai/app/middleware"
)

type RouterGroup struct {
	UserRouter
	CommonRouter
	BlogRouter
}

func (r RouterGroup) InitSystemRouter(router *gin.RouterGroup) {

	// common router
	commonRouter := router.Group("")
	{
		r.InitCommonRouter(commonRouter)
		blogRouter := router.Group("blog")
		{
			r.InitBlogRouter(blogRouter)
		}
	}

	// after interface must with token
	router.Use(middleware.TokenAuthorizeMiddleware())

	userRouter := router.Group("user")
	{
		r.InitUserRouter(userRouter)
		blogRouter := router.Group("blog")
		{
			r.InitBlogPrivateRouter(blogRouter)
		}
	}
}
