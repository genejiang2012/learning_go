package main

import (
	"errors"
	"fmt"
)

func DivideZero(x float64) float64 {
	err := recover("division by zeo")

	if err != nil {
		fmt.Println("The result is not true!")
		return nil
	}
	first_number := 10.0
	second_number := 0.0

	var res = first_number / second_number
	return res
}

func main() {

}
