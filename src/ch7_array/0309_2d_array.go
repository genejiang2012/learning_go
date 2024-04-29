package main

import (
	"fmt"
)

func main() {
	var myArray [4][6]int

	myArray[1][2] = 1
	myArray[2][1] = 2
	myArray[3][1] = 3

	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			fmt.Print(myArray[i][j], " ")
		}
		fmt.Println()
	}

}
