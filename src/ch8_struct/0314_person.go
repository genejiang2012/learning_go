package main

import (
	"fmt"
)

type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int
	slice  []int
	mapOne map[string]string
}

func main() {
	var p1 Person
	fmt.Println(p1)

	if p1.ptr == nil {
		fmt.Println("ok1")
	}

	if p1.slice == nil {
		fmt.Println("ok2")
	}

	if p1.mapOne == nil {
		fmt.Println("ok3")
	}

	p1.slice = make([]int, 10)
	p1.slice[0] = 100

	p1.mapOne = make(map[string]string)
	p1.mapOne["keyOne"] = "tome~"

	fmt.Println(p1)
}
