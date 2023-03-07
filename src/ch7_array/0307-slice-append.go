package main

import (
	"fmt"
)

func main() {
	var slice3 []int = []int{100, 200, 300}

	fmt.Println("The slice before appending ", slice3)

	slice3 = append(slice3, 400, 500, 600)
	fmt.Println("The slice after appedning ", slice3)

}
