package piscine

func Capitalize(s string) string {
	wordStart := true
	result := ""
	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			if wordStart {
				if char >= 'a' && char <= 'z' {
					result += string(char - 32)
				} else {
					result += string(char)
				}
				wordStart = false
			} else {
				if char >= 'A' && char <= 'Z' {
					result += string(char + 32)
				} else {
					result += string(char)
				}
			}
		} else {
			result += string(char)
			wordStart = true
		}
	}
	return result
}
