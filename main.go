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
	tt()

	startHTTPServer()

}

func tt() {
	// a := make([]int64, 0)

	a := make([]int64, 0, 20)

	fmt.Println(cap(a), len(a))

	for i := 0; i < 3; i++ {
		a = append(a, 1)
		fmt.Println(cap(a), len(a))
	}

	fmt.Println("sssss")
	fmt.Println(a)

	ss()
}

// AAA ls
type AAA struct {
	A string
	B int
}

func ss() {

	a := new(AAA)

	a.A = "hel lo"

	fmt.Printf("%+v\n", a)
	fmt.Println(a)

}
