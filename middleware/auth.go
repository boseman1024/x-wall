package middleware

import (
	"github.com/gin-gonic/gin"
	"x-wall/serializer"
	"x-wall/util"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(200, serializer.Response{
				Code: 1000,
				Msg:  "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}
		jwtUtil := util.JwtUtil{}
		if claims := jwtUtil.ParseToekn(token); claims == nil {
			c.JSON(200, serializer.Response{
				Code: 1001,
				Msg:  "携带token无效",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func GetCurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
