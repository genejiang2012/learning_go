/*
The funciton is used to judge the suffix of the ini file
*/

package main

import (
	"errors"
	"fmt"
)

func readConf(name string) (err error) {
	if name == "config.yaml" {
		return nil
	} else {
		return errors.New("fail to read the file")
	}
}

func foo() {
	err := readConf("config.yaml")
	if err != nil {
		panic(err)
	}

	fmt.Println("foo() is going ...")
}

func main() {

	foo()
	fmt.Println("main is going ...")
}
