package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("File name missing")
		return
	}

	if len(args) > 2 {
		fmt.Println("Too many arguments")
		return
	}
	filename := args[1]
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print("Error:", err)
		return
	}
	fmt.Print(string(content))
}
