package main

import "github.com/koron/go-mynumber"

func main() {
	// true
	println(mynumber.ValidateStr("123456789018"))
	// false
	println(mynumber.ValidateStr("123456789010"))
}
