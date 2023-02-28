package main

import (
	"fmt"
)

func DivideZero(first_number float64, second_number float64) float64 {

	defer func() {
		err := recover()

		if err != nil {
			fmt.Println("err=", err)
			fmt.Println("send the mail to ...")
		}
	}()

	res := first_number / second_number

	fmt.Println("The result is", res)
	return res

}

func main() {
	res := DivideZero(10.0, 0.0)
	fmt.Println("The result is", res)

}
