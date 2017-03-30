package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	Usage()
	// tt()
	StartHTTPServer()
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
