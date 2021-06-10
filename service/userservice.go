package service

import (
	"github.com/didi/gendry/scanner"
	"x-wall/db"
	"x-wall/model"
	"x-wall/serializer"
	"x-wall/util"
)

type UserLoginService struct {
	Username string `from:"username" json:"username" binding:"required,min=5,max=30"`
	Password string `from:"password" json:"password" binding:"required,min=6,max=30"`
}
type UserRegisterService struct {
	Username        string `from:"username" json:"username" binding:"required,min=5,max=30"`
	Nickname        string `from:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	Password        string `from:"password" json:"password" binding:"required,min=6,max=30"`
	PasswordConfirm string `from:"password_confirm" json:"password_confirm" binding:"required,min=6,max=30"`
}
type UserService struct{}

func (service *UserLoginService) Login() (string, model.User, *serializer.Response) {
	var user model.User
	rows,err := db.DB.Query("select * from users where username="+service.Username+" limit 1")
	scanner.Scan(rows, &user)
	if  err != nil {
		return "", user, &serializer.Response{
			Code: 102,
			Msg:  "账号或密码错误",
		}
	}
	if isPwd := user.CheckPwd(service.Password); !isPwd {
		return "", user, &serializer.Response{
			Code: 103,
			Msg:  "账号或密码错误",
		}
	}
	jwtUtil := util.JwtUtil{}
	token, err := jwtUtil.CreateToken(user.ID, user.Username)
	if err != nil {
		return "", user, &serializer.Response{
			Code: 104,
			Msg:  "Token生成失败",
		}
	}
	user.Password = ""
	return token, user, nil
}


func (service *UserRegisterService) Valid() *serializer.Response {
	if service.Password != service.PasswordConfirm {
		return &serializer.Response{
			Code: 114,
			Msg:  "两次输入的密码不匹配",
		}
	}
	var user model.User
	rows,_ := db.DB.Query("select * from users where nickname="+service.Nickname+" limit 1")
	scanner.Scan(rows, &user)
	if user.Nickname != "" {
		return &serializer.Response{
			Code: 115,
			Msg:  "当前昵称已被占用",
		}
	}
	rows,_ = db.DB.Query("select * from users where username="+service.Username+" limit 1")
	scanner.Scan(rows, &user)
	if user.Username != "" {
		return &serializer.Response{
			Code: 116,
			Msg:  "当前用户名已被占用",
		}
	}
	return nil
}
