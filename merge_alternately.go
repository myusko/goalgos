package goalgos

import "bytes"

func MergeAlternately(word1 string, word2 string) string {
	var buf bytes.Buffer
	var i int

	buf.Grow(len(word1) + len(word2))

	for len(word1) > i || len(word2) > i {
		if len(word1) > i {
			buf.WriteString(string(word1[i]))
		}
		if len(word2) > i {
			buf.WriteString(string(word2[i]))
		}

		i++
	}

	return buf.String()
}
