package service

import (
	"x-wall/db"
)

type DbConnectService struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

type DbService struct {
	SqlText string `json:"sqlText"`
}

func (dbConnectService *DbConnectService) Connect() (status bool) {
	status = db.Init(dbConnectService.Name, dbConnectService.Username, dbConnectService.Password, dbConnectService.Host, dbConnectService.Port)
	return status
}

func (dbConnectService *DbConnectService) Disconnect() (status bool) {
	status = db.Disconnect()
	return status
}

func (dbService *DbService) ExecSql() (results []map[string]string, err error) {
	rows, err := db.DB.Query(dbService.SqlText)
	if err != nil {
		return nil, err
	}
	//关闭rows链接
	defer rows.Close()
	//读出查询出的列字段名
	cols, _ := rows.Columns()
	//values是每个列的值，这里获取到byte里
	values := make([][]byte, len(cols))
	//query.Scan的参数，因为每次查询出来的列是不定长的，用len(cols)定住当次查询的长度
	scans := make([]interface{}, len(cols))
	//让每一行数据都填充到[][]byte里面,狸猫换太子
	for i := range values {
		scans[i] = &values[i]
	}
	results = make([]map[string]string, 0, 10)
	for rows.Next() {
		err := rows.Scan(scans...)
		if err != nil {
			return nil, err
		}
		row := make(map[string]string, 10)
		for k, v := range values { //每行数据是放在values里面，现在把它挪到row里
			key := cols[k]
			row[key] = string(v)
		}
		results = append(results, row)
	}
	return results, nil
}
