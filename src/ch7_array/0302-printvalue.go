package main

import (
	"fmt"
)

func main() {
	heroes := [...]string{"Spider", "Captain", "Green Gigant"}

	for index, value := range heroes {
		fmt.Printf("The heroes[%d] is %v \n", index, value)
	}

	for _, value := range heroes {
		fmt.Printf("The heroes is %v \n", value)
	}

	var myChar [26]byte

	for i := 0; i < 26; i++ {
		myChar[i] = 'A' + byte(i)
	}

	for i := 0; i < 26; i++ {
		fmt.Printf("%c ", myChar[i])
	}
}
