package main

import (
	"fmt"
	"strconv"
)

// Usage is
func Usage() {

	var welcomeStr = `欢迎使用  微服务 
	- 高性能、跨平台、易扩展 
	- Version: ` + Config.APP.Version + `
	- 服务占用  ` + strconv.Itoa(Config.APP.Port) + `  端口号
	
	请打开浏览器输入 
	http://localhost:` + strconv.Itoa(Config.APP.Port) + `
	`
	fmt.Println(welcomeStr)
}
