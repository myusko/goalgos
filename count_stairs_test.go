package goalgos

import "testing"

func TestCountStairs(t *testing.T) {
	testCases := []struct {
		Input          int
		ExpectedResult int
	}{
		{Input: 8, ExpectedResult: 3},
		{Input: 5, ExpectedResult: 2},
	}

	for _, tc := range testCases {
		actualResult := CountStairs(tc.Input)

		if actualResult != tc.ExpectedResult {
			t.Fatalf("Actual: %v, Want: %v", actualResult, tc.ExpectedResult)
		}
	}
}
