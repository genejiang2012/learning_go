// design one array to store the music character

package main

import (
	"fmt"
)

func main() {
	var arrayCharList [7]string
	arrayCharList[0] = "do"
	arrayCharList[1] = "re"
	arrayCharList[2] = "mi"

	fmt.Println(arrayCharList[0])

	for i := 0; i < 7; i++ {
		fmt.Println(arrayCharList[i])
	}

	for i := 0; i < len(arrayCharList); i++ {
		fmt.Println(i, arrayCharList[i])
	}

	for item, value := range arrayCharList {
		fmt.Println(item, value)
	}

	for _, value := range arrayCharList {
		fmt.Println(value)
	}

	numbers := [3]float64{71.8, 56.2, 89.5}
	var sum float64 = 0.0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println("The total values is ", sum)
}
