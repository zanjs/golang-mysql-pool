package main

// Config is ...
var Config = struct {
	APP struct {
		Port    int    `default:"12300"`
		Version string `default:"0.0.1"`
	}
	DB struct {
		UserName string `default:"root"`
		PassWord string `default:"root"`
		DBName   string `default:"gotest"`
	}
}{}
