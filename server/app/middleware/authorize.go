package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/noobHuKai/app/g"
	"github.com/noobHuKai/app/model/common/response"
	"strconv"
	"strings"
)

var ctx = context.Background()

func TokenAuthorizeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get token
		rawToken := c.Request.Header.Get("Authorization")
		if rawToken == "" {
			response.FailUnauthorized(c, g.ErrTokenNotFound.Error())
			return
		}
		if !strings.HasPrefix(rawToken, "Bearer ") {
			response.FailUnauthorized(c, g.ErrTokenNotFound.Error())
			return
		}
		token := strings.TrimPrefix(rawToken, "Bearer ")

		uidStr, err := g.RDB.Get(ctx, token).Result()
		if err != nil {
			g.Logger.Error(err.Error())
			response.FailUnauthorized(c, g.ErrTokenIsExpired.Error())
			return
		}

		RefreshToken(token)

		uid, err := strconv.Atoi(uidStr)
		if err != nil {
			response.FailUnauthorized(c, g.ErrTokenNotFound.Error())
			return
		}
		c.Set("uid", uint(uid))
		c.Next()
	}
}

func RefreshToken(token string) {
	duration := g.RDB.TTL(ctx, token).Val()
	if duration.Minutes() < g.TimeRefreshToken {
		g.RDB.Expire(ctx, token, g.TimeExpireToken)
	}
}
