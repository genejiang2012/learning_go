package main

import (
	"fmt"
)

func main() {
	var localArray [5]float64

	for i := 0; i < len(localArray); i++ {
		fmt.Printf("Please input the %d value: ", i+1)
		fmt.Scanln(&localArray[i])
	}

	for i := 0; i < len(localArray); i++ {
		fmt.Printf("localArray[%d]=%v \n ", i, localArray[i])
	}

}
