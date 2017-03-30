package main

import (
	"database/sql"

	"github.com/jinzhu/configor"
)

// DB is ...
var DB *sql.DB

func init() {

	configor.Load(&Config, "config.yml")

	DB, _ = sql.Open("mysql", Config.DB.UserName+":"+Config.DB.PassWord+"@/"+Config.DB.DBName+"?charset=utf8")
	DB.SetMaxIdleConns(2000)
	DB.SetMaxIdleConns(1000)
	DB.Ping()
}

// Record is ...
func Record(rows *sql.Rows) (record map[string]string) {

	columns, _ := rows.Columns()

	scanArgs := make([]interface{}, len(columns))

	values := make([]interface{}, len(columns))

	for j := range values {

		scanArgs[j] = &values[j]

	}

	record = make(map[string]string)

	for rows.Next() {

		//将行数据保存到record字典

		err := rows.Scan(scanArgs...)

		CheckErr(err)

		for i, col := range values {

			if col != nil {

				record[columns[i]] = string(col.([]byte))

			}

		}

	}

	return
}

// RSQL is ...
func RSQL(sql string) (record map[string]string) {
	// rows, _ := DB.Query("SELECT * FROM user")
	rows, err := DB.Query(sql)
	defer rows.Close()

	record = Record(rows)

	CheckErr(err)

	return record
}
