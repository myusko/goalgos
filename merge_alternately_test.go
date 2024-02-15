package goalgos

import "testing"

func TestMergeAlternately(t *testing.T) {
	testCases := []struct {
		wordOne        string
		wordTwo        string
		expectedResult string
	}{
		{
			wordOne:        "abc",
			wordTwo:        "pqr",
			expectedResult: "apbqcr",
		},
		{
			wordOne:        "ab",
			wordTwo:        "pqrs",
			expectedResult: "apbqrs",
		},
		{
			wordOne:        "abcd",
			wordTwo:        "pq",
			expectedResult: "apbqcd",
		},
	}

	for _, tc := range testCases {
		actualResult := MergeAlternately(tc.wordOne, tc.wordTwo)

		if actualResult != tc.expectedResult {
			t.Errorf("Got: %v, Want: %v", actualResult, tc.expectedResult)
		}
	}
}
