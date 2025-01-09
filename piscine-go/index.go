package piscine

func Index(s string, toFind string) int {
	for i := 0; i <= len(s)-len(toFind); i++ {
		ismatch := true
		for j := 0; j < len(toFind); j++ {
			if s[i+j] != toFind[j] {
				ismatch = false
				break
			}
		}
		if ismatch {
			return i
		}
	}
	return -1
}
