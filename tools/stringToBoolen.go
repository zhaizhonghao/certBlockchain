package main

import (
	"fmt"
	"strconv"
)

func main() {
	b, err := strconv.ParseBool("true") // string 转bool
	if err == nil {
		fmt.Println(b)
		s := strconv.FormatBool(true) // bool 转string
		fmt.Println(s)
	}
}
