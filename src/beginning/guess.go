package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	/*
		target := rand.Intn(100) + 1
		fmt.Println(target)
	*/
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println(target)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Make a guess: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	guess, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}

	if guess < target {
		fmt.Println("Oops, Your guess warslow.")
	} else if guess > target {
		fmt.Println("Oops, Your guess was high.")
	} else {
		fmt.Print("Bigoo, Your guess is correct.")
	}

}
