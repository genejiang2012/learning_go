package main

import (
	"fmt"
)

func getVariablePara(n1 int, args ...int) int {
	sum := n1

	for i := 0; i < len(args); i++ {
		sum += args[i]
	}

	return sum
}

func main() {
	sum := getVariablePara(10, 20, -20, -10)

	fmt.Println("The result is:", sum)
}
