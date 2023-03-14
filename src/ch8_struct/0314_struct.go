package main

import (
	"fmt"
)

type Cat struct {
	Name  string
	Age   int
	Color string
	Hobby string
}

func main() {
	var catOne Cat
	catOne.Name = "小白"
	catOne.Age = 3
	catOne.Color = "White"
	catOne.Hobby = "eat fishes!"

	fmt.Println("catOne=", catOne)
	fmt.Println("name", catOne.Name)
	fmt.Println("age", catOne.Age)
	fmt.Println("color", catOne.Color)
	fmt.Println("hobby", catOne.Hobby)
}
