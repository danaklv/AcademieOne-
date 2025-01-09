package main

import (
	"bytes"
	"fmt"
	"gomod/functions"
	"os"
)


func main() {
	if len(os.Args[1:]) != 1 {
		fmt.Println("There must be 1 argument")
		os.Exit(0)
	}
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Missing file")
		os.Exit(0)
	}	
	if len(content) == 0 {
	   fmt.Println("no tetromino")
	   os.Exit(0)
	}
	sep := []byte{10} //разделяетс каждую строку (10 - \n)
	SplitContent := bytes.Split(content, sep)
	functions.CheckFormat(SplitContent)
	
	tetrominoesList := functions.FindTetro(SplitContent)
	functions.CheckTetromino(tetrominoesList)

	minSize := functions.FindBoartMinSize(tetrominoesList)
    functions.CreateBoard(minSize)
	functions.TryPosition(0, tetrominoesList, minSize)
}

