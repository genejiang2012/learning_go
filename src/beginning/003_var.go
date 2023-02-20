package main

import (
	"fmt"
)

func main() {
	var intVal int = 10
	var strVal string

	strVal = "The first variables"

	intVal2 := 30

	float32Val3 := 20.7

	// var float32Val4 int
	// float32Val4 = 100.001

	var (
		intVal4 int    = 12
		strVal5 string = "The second variables"
	)

	var intVal6 int32
	var fltVal7 float32

	fmt.Println(intVal, intVal2, strVal, float32Val3, intVal4, strVal5)
	fmt.Print(intVal6, fltVal7)
	fmt.Printf("%T, %T, %T, %T, %T", intVal, intVal2, strVal, intVal6, fltVal7)
}
