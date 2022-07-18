package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		cors.New(cors.Config{
			// 允许所有跨域请求
			AllowAllOrigins: true,
			// 允许所有请求方式
			AllowMethods: []string{"*"},
			AllowHeaders: []string{"Origin"},
			ExposeHeaders: []string{"content-Length"},
			// 域请求持续时间，通过可以保存12小时，持续时间内可以不再进行域请求
			MaxAge: 12 * time.Hour,
		})
	}
}