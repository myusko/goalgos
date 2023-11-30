package goalgos

func SockMerchant(n int32, socks []int32) int32 {
	seen := make(map[int32]int32)
	var result int32

	for _, v := range socks {
		seen[v]++
		if seen[v]%2 == 0 {
			result++
		}
	}

	return result
}
