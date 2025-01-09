package piscine

func AppendRange(min, max int) []int {
	var Arange []int

	for i := min; i < max; i++ {
		Arange = append(Arange, i)
	}
	if max > min {
		return Arange
	}
	return nil
}
