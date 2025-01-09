package piscine

func MakeRange(min, max int) []int {
	if max > min {
		size := max - min
		res := make([]int, size)
		for i := 0; i < size; i++ {
			res[i] = min
			min += 1
		}
		return res
	}
	return nil
}
