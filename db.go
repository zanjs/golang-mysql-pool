package main

import "database/sql"

// DB is ...
var DB *sql.DB

func init() {
	DB, _ = sql.Open("mysql", "root:root@/gotest?charset=utf8")
	DB.SetMaxIdleConns(2000)
	DB.SetMaxIdleConns(1000)
	DB.Ping()
}
