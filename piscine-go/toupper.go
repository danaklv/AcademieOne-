package piscine

func ToUpper(s string) string {
	res := ""
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			char -= 32
		}
		res += string(char)
	}
	return res
}
