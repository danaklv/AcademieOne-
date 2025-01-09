package piscine

func FindNextPrime(n int) int {
	if n <= 1 {
		return 2
	}

	isPrime := func(nb int) bool {
		if nb > 9223372036854775807 {
			return false
		}
		if nb <= 1 {
			return false
		}
		if nb == 2 {
			return true
		}
		if nb%2 == 0 {
			return false
		}
		for i := 3; i*i <= nb; i += 2 {
			if nb%i == 0 {
				return false
			}
		}
		return true
	}

	if isPrime(n) {
		return n
	}

	next := n
	for {
		next++
		if isPrime(next) {
			return next
		}
	}
}
