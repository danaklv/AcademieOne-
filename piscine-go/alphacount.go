package piscine

func AlphaCount(s string) int {
	str := []rune(s)
	count := 0
	for _, char := range str {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			count++
		}
	}
	return count
}
