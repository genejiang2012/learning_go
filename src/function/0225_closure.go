package main

import (
	"fmt"
	_ "strings"
)

func AddUpper() func(int) int {
	n := 10
	str := "hello"

	return func(x int) int {
		n += x
		str += string(36)
		fmt.Println("The string is ", str)
		return n
	}
}

func main() {

	local_func := AddUpper()
	fmt.Println(local_func(1))
	fmt.Println(local_func(2))
	fmt.Println(local_func(3))
}
