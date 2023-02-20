package utils

import (
	"fmt"
)

func Calculator(a float64, b float64, operator byte) float64 {
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
