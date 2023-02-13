package main

import (
	"fmt"
)

func main() {
	x := 0
	x++
	fmt.Println(x)
	x++
	fmt.Println(x)
	x--
	fmt.Println(x)

	for x := 1; x <= 3; x++ {
		fmt.Println(x)
	}
}
