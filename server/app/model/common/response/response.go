package response

import (
	"github.com/gin-gonic/gin"
	"github.com/noobHuKai/app/g"
	"net/http"
)

const (
	CodeOk                  = 200 // 通用成功
	CodeBadRequest          = 400 // 通用失败
	CodeUnauthorized        = 401 // 未授权
	CodeUnProcessableEntity = 422 // 请求格式错误
)

type JsonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type Response = JsonResponse

func ResultJson(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
	c.Abort()
}

func ErrResultJson(c *gin.Context, code int, msg string, data interface{}) {
	g.Logger.Error(msg)
	ResultJson(c, code, msg, data)
}

func Ok(c *gin.Context) {
	ResultJson(c, CodeOk, "", nil)
}
func OkWithMsg(c *gin.Context, msg string) {
	ResultJson(c, CodeOk, msg, nil)
}
func OkWithData(c *gin.Context, data interface{}) {
	ResultJson(c, CodeOk, "", data)
}

func Fail(c *gin.Context) {
	ErrResultJson(c, CodeBadRequest, "", nil)
}
func FailWithMsg(c *gin.Context, msg string) {
	ErrResultJson(c, CodeBadRequest, msg, nil)
}
func FailWithData(c *gin.Context, data interface{}) {
	ErrResultJson(c, CodeBadRequest, "", data)
}

func FailFormatError(c *gin.Context, msg string) {
	ErrResultJson(c, CodeUnProcessableEntity, msg, nil)
}

func FailRequestContentError(c *gin.Context) {
	FailFormatError(c, g.ErrRequestContent.Error())
}

func FailUnauthorized(c *gin.Context, msg string) {
	ErrResultJson(c, CodeUnauthorized, msg, nil)
}
