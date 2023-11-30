package goalgos

import "testing"

func TestTimeDifference(t *testing.T) {
	testCases := []struct {
		N              int
		Times          string
		ExpectedResult int
	}{
		{
			N:              2,
			Times:          "00:00 00:00",
			ExpectedResult: 0,
		},
		{
			N:              5,
			Times:          "22:00 21:00 23:50 23:30 23:55",
			ExpectedResult: 5,
		},
		{
			N:              2,
			Times:          "23:59 00:00",
			ExpectedResult: 1,
		},
	}

	for _, tc := range testCases {
		actualResult := TimeDifference(tc.N, tc.Times)

		if actualResult != tc.ExpectedResult {
			t.Errorf("Got: %v, Want: %v", actualResult, tc.ExpectedResult)
		}
	}
}
