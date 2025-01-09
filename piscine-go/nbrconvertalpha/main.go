package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Atoi_bnr(s string) int {
	arrayStr := []rune(s)
	n := len(s)
	ans := 0
	for i := 0; i < n; i++ {
		ans *= 10
		ans += int(arrayStr[i]) - '0'
	}
	return ans
}

func main() {
	args := os.Args
	if len(args) <= 1 {
		return
	}
	if args[1] == "--upper" {
		for _, nbr := range args[2:] {
			if Atoi_bnr(nbr) < 0 || Atoi_bnr(nbr) > 26 {
				z01.PrintRune(32)
			} else {
				z01.PrintRune(rune(Atoi_bnr(nbr)) + 64)
			}
		}
	} else {
		for _, nbr := range args[1:] {
			if Atoi_bnr(nbr) < 0 || Atoi_bnr(nbr) > 26 {
				z01.PrintRune(32)
			} else {
				z01.PrintRune(rune(Atoi_bnr(nbr)) + 96)
			}
		}
	}
	z01.PrintRune('\n')
}
