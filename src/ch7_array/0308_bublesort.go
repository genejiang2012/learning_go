package main

import (
	"fmt"
)

func bubbleSort(arr *[5]int) {
	fmt.Println("The array before", (*arr))

	tmp := 0

	for i := 0; i < len(*arr); i++ {
		for j := 0; j < len(*arr)-i-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				tmp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = tmp
			}
		}
	}

	fmt.Println("The array after", (*arr))
}

func main() {
	arr := [5]int{10, 5, 13, 9, 3}
	bubbleSort(&arr)

	fmt.Println("The array is", arr)

}
