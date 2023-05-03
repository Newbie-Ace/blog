package system

import (
	"context"

	"github.com/noobHuKai/app/service"
)

type ApiGroup struct {
	UserApi
	CommonApi
	BlogApi
}

var (
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
	blogService = service.ServiceGroupApp.SystemServiceGroup.BlogService
)

var (
	ctx = context.Background()
)
