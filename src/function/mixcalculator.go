package main

import (
	"fmt"
)

func MixCalculator(first_value float64, second_value float64, operator byte) {
	var result float64

	switch operator {
	case '+':
		result = first_value + second_value
	case '-':
		result = second_value - first_value
	case '*':
		result = first_value * second_value
	case '/':
		result = first_value / second_value

	}
}
