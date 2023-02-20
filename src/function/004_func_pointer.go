package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myInt int
	fmt.Println(reflect.TypeOf(&myInt))
	fmt.Println(reflect.TypeOf(myInt))
	var myInt2 *int
	fmt.Println(reflect.TypeOf(myInt2))
}
