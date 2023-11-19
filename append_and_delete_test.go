package goalgos

import "testing"

func TestAppendAndDelete(t *testing.T) {
	testCases := []struct {
		S              string
		T              string
		K              int
		ExpectedResult string
	}{
		{
			S:              "hackerhappy",
			T:              "hackerrank",
			K:              9,
			ExpectedResult: "Yes",
		},
		{
			S:              "aba",
			T:              "aba",
			K:              7,
			ExpectedResult: "Yes",
		},
		{
			S:              "ashley",
			T:              "ash",
			K:              2,
			ExpectedResult: "No",
		},
	}

	for _, tc := range testCases {
		actualResult := AppendAndDelete(tc.S, tc.T, tc.K)

		if actualResult != tc.ExpectedResult {
			t.Errorf("Got: %v, Want: %v", actualResult, tc.ExpectedResult)
		}
	}
}
