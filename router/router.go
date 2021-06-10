package router

import (
	"github.com/gin-gonic/gin"
	"x-wall/api"
	"x-wall/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		//用户登录
		v1.POST("user/login", api.UserLogin)

		v1.POST("wall/sql/exec",api.ExecSql)
		v1.POST("wall/sql/connect",api.Connect)
		v1.GET("wall/sql/disconnect",api.Disconnect)

		//权限访问
		auth := v1.Group("/")
		auth.Use(middleware.AuthRequired())
		{

		}
	}
	return r
}
