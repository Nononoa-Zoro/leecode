package solutions

import "math/big"

func uniquePaths(m, n int) int {
	return int(new(big.Int).Binomial(int64(m+n-2), int64(m-1)).Int64())
}
