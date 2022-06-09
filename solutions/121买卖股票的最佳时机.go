package solutions

func maxProfit(prices []int) int {
	// 最大收入
	maxProfit := 0
	// 记录当前已经遍历过的最小值
	minPrice := prices[0]
	for i := 0; i < len(prices); i++ {
		// 更新最小值
		minPrice = min(minPrice, prices[i])
		// 更新最大利润
		maxProfit = max(maxProfit, prices[i]-minPrice)
	}
	return maxProfit
}
