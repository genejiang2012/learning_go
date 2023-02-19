package main

import (
	"fmt"
)

func foo(bar int) {
	bar += 1
	fmt.Println("The foo(): bar value is", bar)
}

func main() {
	temp := 10
	foo(temp)
	fmt.Println("The main(): temp value is", temp)
}
