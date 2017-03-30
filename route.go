package main

import (
	"net/http"
	"strconv"
)

// StartHTTPServer is router
func StartHTTPServer() {

	http.Handle("/", http.FileServer(http.Dir("./frontend")))

	http.HandleFunc("/pool", Pool)
	http.HandleFunc("/api/users", Users)
	err := http.ListenAndServe(":"+strconv.Itoa(Config.APP.Port), nil)
	CheckErr(err)
}
