package main

import (
	"fmt"
)

func reverse(n int) {
	if n > 2 {
		n--
		reverse(n)
	}
	fmt.Println("The reverse value", n)
}

func main() {
	reverse(4)
}
