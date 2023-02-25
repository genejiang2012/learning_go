package main

import (
	"fmt"
)

var (
	multiply = func(a, b int) int { return a * b }
)

func main() {
	sub := func(n1 int, n2 int) int {
		return n1 - n2
	}

	res1 := sub(20, 10)
	fmt.Println("The results is ", res1)

	res2 := sub(10, 20)
	fmt.Println("The results is ", res2)

	res3 := multiply(20, 30)
	fmt.Println("The result is ", res3)
}
