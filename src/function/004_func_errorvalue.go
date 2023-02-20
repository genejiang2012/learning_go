package main

import (
	"fmt"
	"math"
)

func manyResult() (int, bool, string) {
	return 1, true, "hello"
}

func floatPart(number float64) (integerPart int, fraction float64) {
	wholeNumber := math.Floor(number)
	return int(wholeNumber), number - wholeNumber
}

func main() {
	myInt, myBool, myString := manyResult()
	fmt.Println(myInt, myBool, myString)

	cons, remainder := floatPart(1.26)
	fmt.Println(cons, remainder)
}
