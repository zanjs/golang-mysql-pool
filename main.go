package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func startHTTPServer() {
	http.HandleFunc("/pool", Pool)
	err := http.ListenAndServe(":9090", nil)
	CheckErr(err)
}

func main() {
	fmt.Printf("s")
	startHTTPServer()
}
