package goalgos

func CountStairs(n int) int {
	answers := make([]int, 0)

	for i := 0; i < n; i++ {
		if n >= i+1 {
			answers = append(answers, i+1)
			n = n - (i + 1)
		}
	}

	return answers[len(answers)-1]
}
