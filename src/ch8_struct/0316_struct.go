/*
 * @Author: gene.jiang
 * @Date: 2023-03-16 19:06:07
 * @LastEditors: gene.jiang
 * @LastEditTime: 2023-03-16 20:01:17
 * @Description: copy one struct to another
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

/**
 * @description: main function
 * @return: none
 */
func main() {
	// copy the struct
	var monsterOne Monster
	monsterOne.Name = "Green Giant"
	monsterOne.Age = 500

	var monsterTwo Monster
	monsterTwo = monsterOne
	monsterTwo.Age = 700

	fmt.Println("The monsterOne is ", monsterOne.Name, monsterOne.Age)
	fmt.Println("The monsterTwo is ", monsterTwo.Name, monsterTwo.Age)

	//definiation method 2
	p2 := Monster{Name: "mary", Age: 20}
	fmt.Println("The p2 is ", p2)

	// definiation method 3
	var MonsterThree (*Monster) = new(Monster)
	(*MonsterThree).Name = "Tome"
	MonsterThree.Name = "Tome"
	(*MonsterThree).Age = 20
	MonsterThree.Age = 20
	fmt.Println("The monster three is", *MonsterThree)
	fmt.Println("The monster three new is", MonsterThree)
}
