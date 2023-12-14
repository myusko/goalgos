package goalgos

// Interested gif https://en.wikipedia.org/wiki/Arithmetic_progression#/media/File:Animated_proof_for_the_formula_giving_the_sum_of_the_first_integers_1+2+...+n.gif

func IsArithmeticSequence(numbers []int) bool {
	diff := numbers[0] - numbers[1]

	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i]-numbers[i+1] != diff {
			return false
		}
	}

	return true
}
