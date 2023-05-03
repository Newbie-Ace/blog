package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/noobHuKai/app/api/v1"
	"github.com/noobHuKai/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type CommonRouter struct {
}

func (c *CommonRouter) InitCommonRouter(router *gin.RouterGroup) {
	commonApi := v1.ApiGroupApp.SystemApiGroup.CommonApi

	// 用户登录
	router.POST("user/login", commonApi.Login)
	// 用户注册
	router.POST("user/register", commonApi.Register)

	// swagger
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
