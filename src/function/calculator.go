package main

import (
	"fmt"
)

func calculator(a float64, b float64, operator byte) float64 {
	var result float64

	switch operator {
	case '+':
		result = a + b
	case '-':
		result = a - b
	case '*':
		result = a * b
	case '/':
		if b == 0 {
			fmt.Println("the divider cannot be zero")
			break
		} else {
			result = a / b
		}

	}

	return result
}

func main() {
	a := 10.5
	b := 0.5
	fmt.Println("The result is", calculator(a, b, '+'))
}
