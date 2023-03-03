package main

import (
	"fmt"
)

func main() {
	fmt.Println()

	var intArray = [...]int{-1, 1999, 99, 20, 199}

	sum := 0
	count := 0
	var maxValue int
	maxValue = intArray[0]

	for i := 0; i < len(intArray); i++ {
		if maxValue <= intArray[i] {
			maxValue = intArray[i]
		}
		count = count + i
		sum += intArray[i]
	}

	fmt.Printf("The max values is %d, the average value is %d", maxValue, sum/count)

	fmt.Printf("The max values is %d, the average value is %.2f", maxValue, float64(sum)/float64(len(intArray)))
}
