package system

import (
	"github.com/gin-gonic/gin"
	"github.com/noobHuKai/app/model/common/response"
	"github.com/noobHuKai/app/model/system"
	systemReq "github.com/noobHuKai/app/model/system/request"
)

// CreateType
// @Tags     Blog,Tag
// @Summary  创建博客类型
// @Produce  application/json
// @Param    data  body      systemReq.CreateBlogType     true "类型名称"
// @Success  200   {object}  response.Response{data=system.SysBlogType}
// @Router   /blog/type [post]
func (api *BlogApi) CreateType(c *gin.Context) {
	var req systemReq.CreateBlogType
	if err := c.BindJSON(&req); err != nil {
		response.FailRequestContentError(c)
		return
	}

	blogType := &system.SysBlogType{Name: req.Name}

	blogTypeInter, err := blogService.CreateBlogType(*blogType)
	if err != nil {
		response.FailWithMsg(c, err.Error())
		return
	}

	response.OkWithData(c, blogTypeInter)
}
