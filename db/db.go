package db

import (
	"database/sql"
	"fmt"
	"github.com/didi/gendry/manager"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

var DB *sql.DB

func Init(dbName string, user string, pwd string, host string, portStr string) (status bool) {
	if DB != nil {
		DB.Close()
	}
	//"shares", "root", "123456", "localhost" 3306
	port, _ := strconv.ParseInt(portStr, 10, 32)
	db, err := manager.New(dbName, user, pwd, host).Set(
		manager.SetCharset("utf8"),
		manager.SetAllowCleartextPasswords(true),
		manager.SetInterpolateParams(true),
		manager.SetTimeout(1*time.Second),
		manager.SetReadTimeout(1*time.Second),
	).Port(int(port)).Open(true)

	if err != nil {
		fmt.Println("连接数据库失败")
		return false
	}
	DB = db
	return true
}

func Disconnect() (status bool) {
	if DB != nil {
		DB.Close()
	}
	return true
}
