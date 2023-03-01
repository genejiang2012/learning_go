package main

import (
	"fmt"
)

func main() {
	var numsArray [3]float64 = [3]float64{10.0, 20.0, 30.0}
	fmt.Println("The numsArray is ", numsArray)

	var nums = [3]int{100, 1000, 10000}
	fmt.Println("The nums is ", nums)

	var nums2 = [...]float64{10.0, 20.0, 30.0}
	fmt.Println("The nums2 is ", nums2)

	var string2 = [...]string{1: "Tome", 0: "Jack", 2: "Bob"}
	fmt.Println("The string2 is ", string2)

}
