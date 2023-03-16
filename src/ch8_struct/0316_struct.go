/*
 * @Author: gene.jiang
 * @Date: 2023-03-16 19:06:07
 * @LastEditors: gene.jiang
 * @LastEditTime: 2023-03-16 19:19:42
 * @Description: copy one struct to another struct
 * @FilePath: \learning_go\src\ch8_struct\0316_struct.go
 */

package main

import (
	"fmt"
)

type Monster struct {
	Name string
	Age  int
}

func main() {
	var monsterOne Monster
	monsterOne.Name = "Green Giant"
	monsterOne.Age = 500

	var monsterTwo Monster
	monsterTwo = monsterOne
	monsterTwo.Age = 700

	fmt.Println("The monsterOne is ", monsterOne.Name, monsterOne.Age)
	fmt.Println("The monsterTwo is ", monsterTwo.Name, monsterTwo.Age)
}
