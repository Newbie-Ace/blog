package g

import "errors"

var (
	ErrUserNotFound      = errors.New("用户不存在")
	ErrUserPassWordWrong = errors.New("用户密码错误")
	ErrUserAlreadyExist  = errors.New("用户已存在")

	ErrTokenNotFound  = errors.New("未找到登录信息")
	ErrTokenIsExpired = errors.New("登录已失效")

	ErrRequestContent = errors.New("错误请求内容")
)
