/*
 * @Author: gene.jiang
 * @Date: 2023-03-17 19:17:59
 * @LastEditors: gene.jiang
 * @LastEditTime: 2023-03-17 19:21:58
 * @Description: struct example
 * @FilePath: \learning_go\src\ch8_struct\0317_struct_example.go
 */

package main

import (
	"fmt"
)

var aStruct struct {
	shortName   int
	longerName  float64
	longestName string
}

func main() {
	aStruct.shortName = 200
	aStruct.longerName = 3.14
	aStruct.longestName = "Bob Jiang"

	fmt.Println(aStruct.shortName)
	fmt.Println(aStruct.longerName)
	fmt.Println(aStruct.longestName)
}
