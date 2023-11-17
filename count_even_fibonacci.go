package goalgos

func CountEvenFibonacci(n uint64) uint64 {
	var a, b, result uint64
	b = 1
	copyN := n

	for n--; n > 0; n-- {
		a += b
		a, b = b, a

		if b > copyN {
			return result
		}

		if b%2 == 0 {
			result += b
		}
	}

	return result
}
