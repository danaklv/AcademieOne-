package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	progName := os.Args[0]
	progNameRunes := []rune(progName)
	if len(progNameRunes) >= 2 && progNameRunes[0] == '.' && progNameRunes[1] == '/' {
		progNameRunes = progNameRunes[2:]
	}
	for _, char := range progNameRunes {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
