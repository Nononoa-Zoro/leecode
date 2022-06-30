package solutions

func coinChange(coins []int, amount int) int {
	if len(coins) == 0 {
		return -1
	}
	dp := make([]int, amount+1, amount+1)
	// 初始化为最大的元素
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			// 当前面额大于可用的小面额
			if i >= coins[j] {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
