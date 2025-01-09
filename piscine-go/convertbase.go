package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	decimalValue := 0
	sourceBase := len(baseFrom)
	targetBase := len(baseTo)

	for _, digit := range nbr {
		sourceDigitValue := 0
		for j, char := range baseFrom {
			if char == digit {
				sourceDigitValue = j
				break
			}
		}
		decimalValue = decimalValue*sourceBase + sourceDigitValue
	}

	var result []rune
	for decimalValue > 0 {
		remainder := decimalValue % targetBase
		result = append([]rune{rune(baseTo[remainder])}, result...)
		decimalValue /= targetBase
	}

	if len(result) == 0 {
		result = []rune{rune(baseTo[0])}
	}
	return string(result)
}
