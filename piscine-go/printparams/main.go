package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	for _, arg := range args {
		for i := 0; i < len(arg); i++ {
			z01.PrintRune(rune(arg[i]))
		}
		z01.PrintRune('\n')
	}
}
