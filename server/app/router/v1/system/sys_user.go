package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/noobHuKai/app/api/v1"
)

type UserRouter struct {
}

func (u *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	userApi := v1.ApiGroupApp.SystemApiGroup.UserApi
	{
		// 获取用户信息
		router.GET("", userApi.Info)
	}
}
