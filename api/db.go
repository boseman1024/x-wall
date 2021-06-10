package api

import (
	"github.com/gin-gonic/gin"
	"x-wall/serializer"
	"x-wall/service"
)

func ExecSql(c *gin.Context){
	var dbService service.DbService
	if err := c.ShouldBind(&dbService); err == nil {
		if result, err := dbService.ExecSql(); err != nil {
			c.JSON(200, err)
		} else {
			c.JSON(200, serializer.Response{
				Code: 200,
				Data: result,
			})
		}
	} else {
		c.JSON(200, serializer.Response{
			Code: 111,
			Msg:  "sql服务初始化失败",
		})
	}
}

func Connect(c *gin.Context){
	var dbConnectService service.DbConnectService
	if err := c.ShouldBind(&dbConnectService); err == nil {
		status := dbConnectService.Connect()
		c.JSON(200, serializer.Response{
			Code: 200,
			Data: status,
		})
	} else {
		c.JSON(200, serializer.Response{
			Code: 111,
			Msg:  "sql服务初始化失败",
		})
	}
}

func Disconnect(c *gin.Context){
	var dbConnectService service.DbConnectService
	status := dbConnectService.Disconnect()
	c.JSON(200, serializer.Response{
		Code: 200,
		Data: status,
	})
}
