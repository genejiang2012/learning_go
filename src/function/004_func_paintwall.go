package main

import (
	"fmt"
)

func calcWallArea(width float64, height float64) float64 {
	return width * height / 10
}

func main() {
	var amount, total float64
	amount = calcWallArea(10, 20)
	fmt.Printf("%0.2f liters needs\n", amount)
	total += amount
	amount = calcWallArea(20, 30)
	fmt.Printf("%0.2f liters needs\n", amount)
	total += amount
	fmt.Printf("%0.2f total liters needs\n", total)
}
