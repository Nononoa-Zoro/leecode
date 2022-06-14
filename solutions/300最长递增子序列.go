package solutions

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 1
	}
	// dp[i]表示以i结尾的最长递增子序列长度
	dp := make([]int, n, n)
	dp[0] = 1
	var res int
	for i := 1; i < n; i++ {
		var maxBeforeI int
		for j := 0; j < i; j++ {
			// 找到i之前的最大的递增子序列长度
			if nums[j] < nums[i] {
				maxBeforeI = max(maxBeforeI, dp[j])
			}
		}
		dp[i] = maxBeforeI + 1
		res = max(res, dp[i])
	}
	return res
}
