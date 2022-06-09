package solutions

func maxAbsoluteSum(nums []int) int {
	prefix := make([]int, 0, len(nums))
	sum := 0
	for _, v := range nums {
		sum += v
		prefix = append(prefix, sum)
	}

	// 前缀和负数最小值
	minProfit := min(0, prefix[0])
	// 前缀和正数最大值
	maxProfit := max(0, prefix[0])
	for _, v := range prefix {
		if v < minProfit {
			minProfit = v
		}
		if v > maxProfit {
			maxProfit = v
		}
	}

	return maxProfit - minProfit
}
