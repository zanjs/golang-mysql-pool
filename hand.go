package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

// Pool is ...
func Pool(w http.ResponseWriter, r *http.Request) {
	users := ModelUser()

	// 类型打印
	fmt.Println(reflect.TypeOf(users))
	// map 返回json 编码
	str, err := json.Marshal(users)

	CheckErr(err)

	fmt.Fprintln(w, string(str))
}

// Users is ...
func Users(w http.ResponseWriter, r *http.Request) {
	users := ModelUser()

	str, err := json.Marshal(users)

	CheckErr(err)

	fmt.Fprintln(w, string(str))

}
