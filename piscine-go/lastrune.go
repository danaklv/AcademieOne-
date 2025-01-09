package piscine

func LastRune(s string) rune {
	str := []rune(s)
	return str[len(s)-1]
}
