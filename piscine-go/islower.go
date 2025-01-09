package piscine

func IsLower(s string) bool {
	count := 0
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			count++
		}
	}
	if count == len(s) {
		return true
	}
	return false
}
