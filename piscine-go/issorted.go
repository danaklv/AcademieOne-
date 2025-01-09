package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	length := len(tab)
	cap := length
	sort := 0

	for i := 0; i < length-1; i++ {
		isSorted := f(tab[i], tab[i+1])
		if isSorted > 0 {
			sort++
		} else if isSorted < 0 {
			sort--
		} else {
			cap--
		}
	}
	cap--
	return sort == cap || sort == -cap
}
