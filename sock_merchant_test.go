package goalgos

import "testing"

func TestSockMerchant(t *testing.T) {
	testCases := []struct {
		N              int32
		Socks          []int32
		ExpectedResult int32
	}{
		{
			N:              int32(5),
			Socks:          []int32{1, 1, 2, 2, 2},
			ExpectedResult: 2,
		},
		{
			N:              int32(8),
			Socks:          []int32{1, 1, 2, 2, 2, 2},
			ExpectedResult: 3,
		},
		{
			N:              int32(4),
			Socks:          []int32{1, 2, 3, 4},
			ExpectedResult: 0,
		},
	}

	for _, tc := range testCases {
		actualResult := SockMerchant(tc.N, tc.Socks)

		if actualResult != tc.ExpectedResult {
			t.Errorf("Got: %v, Want: %v", actualResult, tc.ExpectedResult)
		}
	}
}
