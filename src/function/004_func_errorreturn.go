package main

import (
	// "errors"
	"fmt"
)

func printNeeds(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("the width is invalid")
	}

	if height < 0 {
		return 0, fmt.Errorf("the height is invalid")
	}

	return width * height / 10, nil
}

func main() {
	// error := errors.New("The is new error values")
	// fmt.Println(error)

	amount, err := printNeeds(4.0, 2.0)
	fmt.Println(err)
	fmt.Printf("The amount is %f", amount)
}
