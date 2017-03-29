package main

import (
	"database/sql"
	"fmt"
	"reflect"
)

// User is ...
func User() (rows *sql.Rows, record map[string]string) {
	rows, err := DB.Query("SELECT * FROM user")
	defer rows.Close()

	CheckErr(err)

	columns, _ := rows.Columns()

	scanArgs := make([]interface{}, len(columns))

	values := make([]interface{}, len(columns))

	for j := range values {

		scanArgs[j] = &values[j]

	}

	record = make(map[string]string)

	for rows.Next() {

		//将行数据保存到record字典

		err = rows.Scan(scanArgs...)

		for i, col := range values {

			if col != nil {

				record[columns[i]] = string(col.([]byte))

			}

		}

	}

	fmt.Println(reflect.TypeOf(record))

	fmt.Println(record)

	return nil, record
}
