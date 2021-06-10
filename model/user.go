package model

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint    `json:id`
	Username string  `json:"username"`
	Password string  `json:"password,omitempty"`
	Nickname string  `json:"nickname"`
	Status   string  `json:"status,omitempty"`
	Avatar   string  `json:"avatar"`
	Note     string  `json:"note"`
}

const (
	//加密难度
	PasswordCost = 12
)

func (user *User) SetPwd(pwd string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), PasswordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPwd(pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pwd))
	return err == nil
}
