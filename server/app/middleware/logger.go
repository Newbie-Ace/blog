package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/noobHuKai/app/g"
	"github.com/noobHuKai/app/utils"
	"io"
	"os"
	"path"
	"strings"
)

func LoggerConfigMiddleware() gin.HandlerFunc {
	filename := "logs/access/access"
	fileWriter := utils.GetRotateLogWriter(filename+"_%Y%m%d.log", fmt.Sprintf("logs/%s.log", path.Base(filename)))
	loggerWriter := io.MultiWriter(fileWriter, os.Stdout)

	loggerFormat := func(param gin.LogFormatterParams) string {
		if strings.HasPrefix(param.Path, "/api/v1/swagger") {
			return ""
		}
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(g.TimeFormat),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}
	return gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: loggerFormat,
		Output:    loggerWriter,
	})
}
