package main

import "fmt"

// CheckErr is ...
func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
