package piscine

func Split(s, sep string) []string {
	var result []string
	start := 0

	for i := 0; i <= len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			if i > start {
				result = append(result, s[start:i])
			}
			start = i + len(sep)
			i = start - 1
		}
	}

	if start < len(s) {
		result = append(result, s[start:])
	}

	return result
}
