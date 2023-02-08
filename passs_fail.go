/*
Requirement Description: read one number from keyboard and judge it whether it passes or fails
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please input one number: ")

	//read one number by keyboard
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	// change the input from the string to number
	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}

	// judge the number
	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}

	fmt.Println("The finial status is ", status)

}
