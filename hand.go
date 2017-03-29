package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

// Pool is ...
func Pool(w http.ResponseWriter, r *http.Request) {
	rows, record := User()

	// 类型打印
	fmt.Println(reflect.TypeOf(rows))
	fmt.Println(reflect.TypeOf(record))
	// map 返回json 编码
	str, err := json.Marshal(record)

	CheckErr(err)

	fmt.Fprintln(w, string(str))
}
