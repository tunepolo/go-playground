package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var filename = "./fileread.go"

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("could not open:" + filename)
	}

	defer file.Close()

	contents, err := ioutil.ReadAll(file)
	fmt.Println(string(contents))
}
