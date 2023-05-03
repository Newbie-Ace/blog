package main

import (
	"github.com/noobHuKai/app/initialize"
)

// @title           Blog API
// @version         1.0

// @contact.name   noobHuKai
// @contact.url    https://github.com/noobHuKai

// @license.name  MIT
// @license.url   https://opensource.org/license/mit/

// @host      localhost:8000
// @BasePath  /api/v1

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	initialize.InitServer()
	initialize.RunServer()
}
