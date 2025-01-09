package piscine

func Rot14(s string) string {
	res := ""
	for _, r := range s {
		if (r >= 65 && r <= 76) || (r >= 97 && r <= 108) {
			r = r + 14
		} else if (r >= 77 && r <= 90) || (r >= 109 && r <= 122) {
			r = r - 12
		}
		res += string(r)
	}
	return res
}
