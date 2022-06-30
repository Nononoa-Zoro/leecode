package solutions

import "sort"

// https://leetcode.cn/problems/3sum-closest/
func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	ans := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		L, R := i+1, len(nums)-1
		num := nums[i] + nums[L] + nums[R]
		if abs(target, num) < abs(target, ans) {
			ans = num
		}
		if num < target {
			L++
		} else if num > target {
			R--
		} else {
			return ans
		}
	}
	return ans
}

// abs 求两个数的绝对值
func abs(i, j int) int {
	if i > j {
		return i - j
	}
	return j - i
}
