package goalgos

import "testing"

func TestIsArithmeticSequence(t *testing.T) {
	testCases := []struct {
		Input          []int
		ExpectedResult bool
	}{
		{
			Input:          []int{2, 4, 6, 8, 10, 12, 14, 16, 18},
			ExpectedResult: true,
		},
		{
			Input:          []int{1, 5, 9, 13, 17, 21, 25, 29, 33},
			ExpectedResult: true,
		},
		{
			Input:          []int{1, 5, 9, 13, 17, 21, 25, 29, 34},
			ExpectedResult: false,
		},
		{
			Input:          []int{2, 4, 6, 8, 10, 12, 14, 16, 20},
			ExpectedResult: false,
		},
		{
			Input:          []int{2, 5, 10},
			ExpectedResult: false,
		},
	}

	for _, tc := range testCases {
		actualResult := IsArithmeticSequence(tc.Input)

		if actualResult != tc.ExpectedResult {
			t.Errorf("Got: %v, Want: %v", actualResult, tc.ExpectedResult)
		}
	}
}
