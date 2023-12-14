package goalgos

func IsArithmeticSequence(numbers []int) bool {
	diff := numbers[0] - numbers[1]

	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i]-numbers[i+1] != diff {
			return false
		}
	}

	return true
}
