package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/noobHuKai/app/api/v1"
)

type BlogRouter struct {
}

func (r *BlogRouter) InitBlogRouter(router *gin.RouterGroup) {
	blogApi := v1.ApiGroupApp.SystemApiGroup.BlogApi

	// 获取博客列表
	router.GET("", blogApi.ListBlog)

	// 获取博客内容
	router.GET(":id", blogApi.BlogContent)

}

func (r *BlogRouter) InitBlogPrivateRouter(router *gin.RouterGroup) {
	blogApi := v1.ApiGroupApp.SystemApiGroup.BlogApi

	// 创建博客标签
	router.POST("tag", blogApi.CreateTag)

	// 创建博客类型
	router.POST("type", blogApi.CreateType)

	// 创建博客
	router.POST("", blogApi.CreateBlog)
}
