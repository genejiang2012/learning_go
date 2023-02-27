package main

import (
	"fmt"
)

var age int = 50
var Name string = "Tom~"

func test() {
	age := 10
	Name := "Jack~"
	fmt.Println("age=", age)
	fmt.Println("Name=", Name)
}

func main() {
	fmt.Println("age=", age)
	fmt.Println("Name=", Name)
	test()
}
