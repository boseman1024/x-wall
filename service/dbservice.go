package service

import (
	"github.com/didi/gendry/scanner"
	"x-wall/db"
)

type DbConnectService struct {
	Name string `json:"name"`
	Username   string `json:"username"`
	Password    string `json:"password"`
	Host   string `json:"host"`
	Port   string `json:"port"`
}

type DbService struct {
	SqlText string `json:"sqlText"`
}

func (dbConnectService *DbConnectService) Connect() (status bool) {
	status = db.Init(dbConnectService.Name,dbConnectService.Username,dbConnectService.Password,dbConnectService.Host,dbConnectService.Port)
	return status
}

func (dbConnectService *DbConnectService) Disconnect() (status bool) {
	status = db.Disconnect()
	return status
}

func (dbService *DbService) ExecSql() (result []map[string]interface{},err error) {
	rows,err := db.DB.Query(dbService.SqlText)
	result,err = scanner.ScanMapClose(rows)
	return result,err
}
