package main

import (
	"fmt"
)

func getSumAndSub(first_value int, second_value int) (int, int) {
	sum := first_value + second_value
	sub := first_value - second_value

	return sum, sub
}

func main() {
	sumRes, subRes := getSumAndSub(10, 20)

	fmt.Println("The sum result is", sumRes, "the sub result is", subRes)
}
