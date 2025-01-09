package piscine

func Compact(ptr *[]string) int {
	slice := *ptr
	i := 0
	for _, element := range slice {
		if element != "" {
			slice[i] = element
			i++
		}
	}
	*ptr = slice[:i]
	return i
}
