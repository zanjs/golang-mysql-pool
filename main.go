package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:root@/gotest?charset=utf8")
	db.SetMaxIdleConns(2000)
	db.SetMaxIdleConns(1000)
	db.Ping()
}

func pool(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM user")

	defer rows.Close()

	checkErr(err)

	columns, _ := rows.Columns()

	scanArgs := make([]interface{}, len(columns))

	values := make([]interface{}, len(columns))

	for j := range values {

		scanArgs[j] = &values[j]

	}

	record := make(map[string]string)

	for rows.Next() {

		//将行数据保存到record字典

		err = rows.Scan(scanArgs...)

		for i, col := range values {

			if col != nil {

				record[columns[i]] = string(col.([]byte))

			}

		}

	}

	fmt.Println(record)

	fmt.Fprintln(w, "finish")
}

func startHTTPServer() {
	http.HandleFunc("/pool", pool)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {

	}

}

func main() {
	fmt.Printf("s")
	startHTTPServer()
}

func checkErr(err error) {

	if err != nil {

		fmt.Println(err)

		panic(err)

	}

}
