package main

import (
	"fmt"
)

func bar() {
	fmt.Println("The bar func is executed...")
	panic(-1)
}

func foo() {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println("The foo func is executed...")
		}
	}()

	bar()

}

func main() {
	foo()
	fmt.Println("The main func is executed...")
}
