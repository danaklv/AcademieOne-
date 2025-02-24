package piscine

func IterativePower(nb int, power int) int {
	res := nb
	if power < 0 {
		return 0
	}
	if power == 0 || nb == 1 {
		return 1
	}
	for i := 1; i < power; i++ {
		res *= nb
	}
	return res
}
