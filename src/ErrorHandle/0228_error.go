package main

import (
	"errors"
	"fmt"
)

func readIni(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("读取文件错误。。。")
	}
}

func main() {
	readIni("test.mp4")
	fmt.Println("read the files is correct!")
}
