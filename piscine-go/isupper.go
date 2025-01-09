package piscine

func IsUpper(s string) bool {
	count := 0
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			count++
		}
	}
	if count == len(s) {
		return true
	}
	return false
}
