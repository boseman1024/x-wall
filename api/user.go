package api

import (
	"github.com/gin-gonic/gin"
	"x-wall/serializer"
	"x-wall/service"
)

//注册
func UserRegister(c *gin.Context) {

}

//登录
func UserLogin(c *gin.Context) {
	var userLoginService service.UserLoginService
	if err := c.ShouldBind(&userLoginService); err == nil {
		if token, user, err := userLoginService.Login(); err != nil {
			c.JSON(200, err)
		} else {
			c.JSON(200, serializer.Response{
				Code: 200,
				Data: serializer.LoginResponse{
					Token: token,
					User:  user,
				},
			})
		}
	} else {
		c.JSON(200, serializer.Response{
			Code: 111,
			Msg:  "登录服务初始化失败",
		})
	}
}

func UserLogout(c *gin.Context) {

}

func UserGet(c *gin.Context) {

}

func UserUpdate(c *gin.Context) {

}
