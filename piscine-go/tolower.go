package piscine

func ToLower(s string) string {
	res := ""
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			char += 32
		}
		res += string(char)
	}
	return res
}
