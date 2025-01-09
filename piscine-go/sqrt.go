package piscine

func Sqrt(nb int) int {
	if nb == 0 {
		return 0
	}
	sqrt := 0
	for i := 0; i*i <= nb; i++ {
		sqrt = i
	}
	if sqrt*sqrt == nb {
		return nb / sqrt
	}
	return 0
}
