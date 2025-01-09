package piscine

func SplitWhiteSpaces(s string) []string {
	var result []string
	currentWord := ""
	for _, char := range s {
		if char == ' ' || char == '\t' || char == '\n' {
			if currentWord != "" {
				result = append(result, currentWord)
				currentWord = ""
			}
		} else {
			currentWord += string(char)
		}
	}
	if currentWord != "" {
		result = append(result, currentWord)
	}
	return result
}
