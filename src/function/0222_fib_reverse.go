package main

import (
	"fmt"
)

func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func recursive2(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*recursive2(n-1) + 1
	}
}

func UsePointer(value *int) {
	*value = *value + 10
	fmt.Println("The value in the func", *value)
}

func main() {
	fmt.Println("The result is", fib(2))
	fmt.Println("The result is", fib(3))
	fmt.Println("The result is", fib(4))

	fmt.Println("The func recursive result is", recursive2(1))
	fmt.Println("The func recursive result is", recursive2(2))
	fmt.Println("The func recursive result is", recursive2(3))

	//pointer
	number := 20
	fmt.Println("The nuber before calling the function", number)
	UsePointer(&number)
	fmt.Println("The nuber after calling the function", number)
}
