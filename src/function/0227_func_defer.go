package main

import (
	"fmt"
)

func test(n1 int, n2 int) int {
	defer fmt.Println("The n1 is", n1)
	defer fmt.Println("The n2 is", n2)

	n1++
	n2++

	var res = n1 + n2
	fmt.Println("The res of test is", res)
	return res
}

func main() {
	n1 := 10
	n2 := 20
	res := test(n1, n2)
	fmt.Println("The result of main is", res)
}
