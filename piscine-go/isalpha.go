package piscine

func IsAlpha(s string) bool {
	str := []rune(s)
	count := 0
	for _, char := range str {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			count++
		}
	}
	if count == len(s) {
		return true
	}
	return false
}
