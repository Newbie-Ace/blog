package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/noobHuKai/app/model/common/request"
	"github.com/noobHuKai/app/model/common/response"
	"github.com/noobHuKai/app/model/system"
	systemReq "github.com/noobHuKai/app/model/system/request"
)

type BlogApi struct{}

// CreateBlog
// @Tags     Blog
// @Summary  创建博客
// @Produce  application/json
// @Param    data  body      systemReq.CreateBlog     true "创建博客"
// @Success  200   {object}  response.Response{}
// @Router   /blog [post]
func (api *BlogApi) CreateBlog(c *gin.Context) {

	uid := c.GetUint("uid")
	fmt.Println(uid)

	var req systemReq.CreateBlog
	if err := c.BindJSON(&req); err != nil {
		response.FailRequestContentError(c)
		return
	}

	blog := &system.SysBlog{UserID: uid, Title: req.Title, Content: req.Content, TypeID: req.TypeID}

	err := blogService.CreateBlog(*blog, req.Tags)
	if err != nil {
		response.FailWithMsg(c, err.Error())
		return
	}

	response.OkWithData(c, nil)
}

// ListBlog
// @Tags     Blog
// @Summary  获取博客列表
// @Param    q    query     request.Page  true "分页"
// @Success  200   {object}  response.Response{data=[]system.BlogContent}
// @Router   /blog [get]
func (api *BlogApi) ListBlog(c *gin.Context) {

	var req request.Page
	if err := c.ShouldBind(&req); err != nil {
		response.FailRequestContentError(c)
		return
	}

	res, err := blogService.ListBlog(req)
	if err != nil {
		response.FailWithMsg(c, err.Error())
		return
	}

	response.OkWithData(c, res)
}

// BlogContent
// @Tags     Blog
// @Summary  获取博客内容
// @Param    id   path      int  true  "Blog ID"
// @Success  200   {object}  response.Response{data=system.BlogContent}
// @Router   /blog/{id} [get]
func (api *BlogApi) BlogContent(c *gin.Context) {

	var req request.ID

	if err := c.ShouldBindUri(&req); err != nil {
		response.FailRequestContentError(c)
		return
	}

	res, err := blogService.BlogContent(req.ID)
	if err != nil {
		response.FailWithMsg(c, err.Error())
		return
	}

	response.OkWithData(c, res)
}
