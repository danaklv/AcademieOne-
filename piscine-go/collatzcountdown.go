package piscine

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}
	count := 0
	for i := start; i != 1; {
		if i%2 == 0 {
			i = i / 2
		} else if i%2 != 0 {
			i = 3*i + 1
		}
		count++
	}
	return count
}
