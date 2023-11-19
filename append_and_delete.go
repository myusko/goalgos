package goalgos

func AppendAndDelete(s, t string, k int) string {
	var indexS, indexT, stepsToDo int

	for indexS < len(s) && indexT < len(t) {
		if s[indexS] == t[indexT] {
			indexS++
			indexT++
		} else {
			break
		}
	}

	stepsToDo = (len(s) - indexS) + (len(t) - indexT)

	if stepsToDo > k {
		return "No"
	}

	if k > len(s)+len(t) {
		return "Yes"
	}

	if (k-stepsToDo)%2 == 0 {
		return "Yes"
	}

	return "No"
}
