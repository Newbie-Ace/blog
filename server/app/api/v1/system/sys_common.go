package system

import (
	"github.com/gin-gonic/gin"
	"github.com/noobHuKai/app/g"
	"github.com/noobHuKai/app/model/common/response"
	"github.com/noobHuKai/app/model/system"
	systemReq "github.com/noobHuKai/app/model/system/request"
	systemRes "github.com/noobHuKai/app/model/system/response"
	"github.com/noobHuKai/app/utils"
	uuid "github.com/satori/go.uuid"
)

type CommonApi struct{}

// Login
// @Tags     User
// @Summary  用户登录
// @Produce  application/json
// @Param    data  body      systemReq.Login     true "用户名, 密码"
// @Success  200   {object}  response.Response{data=systemRes.Login}
// @Router   /user/login [post]
func (api *CommonApi) Login(c *gin.Context) {
	var req systemReq.Login
	if err := c.BindJSON(&req); err != nil {
		response.FailRequestContentError(c)
		return
	}
	userInter, err := userService.Login(req.Username, req.Password)
	if err != nil {
		response.FailWithMsg(c, err.Error())
		return
	}
	// generate token
	token := utils.SHA256Encrypt(uuid.NewV4().Bytes())
	// set redis
	g.RDB.Set(ctx, token, userInter.ID, g.TimeExpireToken)
	// response
	res := systemRes.Login{Token: token, ExpiresAt: g.TimeExpireToken.String()}
	response.OkWithData(c, res)
}

// Register
// @Tags     User
// @Summary  用户注册
// @Produce  application/json
// @Param    data  body      systemReq.Register     true "用户名, 密码, 昵称"
// @Success  200   {object}  response.Response{data=system.SysUser}
// @Router   /user/register [post]
func (api *CommonApi) Register(c *gin.Context) {
	var req systemReq.Register
	if err := c.BindJSON(&req); err != nil {
		response.FailRequestContentError(c)
		return
	}

	user := &system.SysUser{Username: req.Username, NickName: req.NickName, Password: req.Password}

	userInter, err := userService.Register(*user)
	if err != nil {
		response.FailWithMsg(c, err.Error())
		return
	}

	response.OkWithData(c, userInter)
}
