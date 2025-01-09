package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	var digits [10]int
	for n > 0 {
		digit := n % 10
		digits[digit]++
		n /= 10
	}
	for digit, count := range digits {
		for i := 0; i < count; i++ {
			z01.PrintRune(rune(digit + '0'))
		}
	}
}
