package goalgos

import "testing"

func TestCountEvenFibonacci(t *testing.T) {
	testCases := []struct {
		Input          uint64
		ExpectedResult uint64
	}{
		{Input: uint64(10), ExpectedResult: uint64(10)},
		{Input: uint64(100), ExpectedResult: uint64(44)},
	}

	for _, tc := range testCases {
		actualResult := CountEvenFibonacci(tc.Input)

		if actualResult != tc.ExpectedResult {
			t.Fatalf("Actual: %v, Want: %v", actualResult, tc.ExpectedResult)
		}
	}
}
