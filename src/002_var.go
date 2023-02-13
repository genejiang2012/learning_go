package main

import (
	"fmt"
)

var packageVar string = "package variable"

func main() {
	var funcVar string = "function variable"

	{
		var blockVar string = "block variable"
		fmt.Println(packageVar, funcVar, blockVar)

	}

	fmt.Println(packageVar, funcVar)
}
