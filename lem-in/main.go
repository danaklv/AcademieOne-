package main

import (
	"fmt"
	utils "lm/utils/check"
	"os"
)


func main() {
	if len(os.Args) != 2 {
		fmt.Println("Incorrect number of arguments")
		return
	}
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
	text := string(content)
	utils.CheckFormat(text)

}
