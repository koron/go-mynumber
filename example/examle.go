package main

import (
	"fmt"

	"github.com/koron/go-mynumber"
)

func main() {
	// true
	fmt.Println(mynumber.ValidateStr("123456789018"))
	// false
	fmt.Println(mynumber.ValidateStr("123456789010"))
}
