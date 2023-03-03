package main

import (
	"fmt"
)

func main() {
	localArray := [...]int{11, 22, 33, 44, 55, 66}

	sliceArray := localArray[2:6]
	fmt.Println("The slice array is", sliceArray)
	fmt.Println("The lenght of the slice array is", len(sliceArray))
	fmt.Println("The cap of the slice array is ", cap(sliceArray))
}
