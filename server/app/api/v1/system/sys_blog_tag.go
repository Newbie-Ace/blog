package system

import (
	"github.com/gin-gonic/gin"
	"github.com/noobHuKai/app/model/common/response"
	"github.com/noobHuKai/app/model/system"
	systemReq "github.com/noobHuKai/app/model/system/request"
)

// CreateTag
// @Tags     Blog,Tag
// @Summary  创建博客标签
// @Produce  application/json
// @Param    data  body      systemReq.CreateBlogTag     true "标签名称"
// @Success  200   {object}  response.Response{data=system.SysTag}
// @Router   /blog/tag [post]
func (api *BlogApi) CreateTag(c *gin.Context) {
	var req systemReq.CreateBlogTag
	if err := c.BindJSON(&req); err != nil {
		response.FailRequestContentError(c)
		return
	}

	tag := &system.SysTag{Name: req.Name}

	tagInter, err := blogService.CreateBlogTag(*tag)
	if err != nil {
		response.FailWithMsg(c, err.Error())
		return
	}

	response.OkWithData(c, tagInter)
}
