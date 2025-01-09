package piscine

func IsNumeric(s string) bool {
	str := []rune(s)
	count := 0
	for _, char := range str {
		if char >= '0' && char <= '9' {
			count++
		}
	}
	if count == len(s) {
		return true
	}
	return false
}
