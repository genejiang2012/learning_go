package main

import (
	"fmt"
)

func main() {
	var slice []float64 = make([]float64, 5, 10)
	slice[0] = 20.0
	slice[3] = 30.0

	for i := 0; i < len(slice); i++ {
		fmt.Printf("The slice array is, %d,  %.2f \n", i, slice[i])
	}

	for index, value := range slice {
		fmt.Printf("The index of the slice is %d, the value of slice is %.2f \n", index, value)
	}
}
