package solutions

import "sort"

func minIncrementForUnique(nums []int) int {
	sort.Ints(nums)
	moved := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			pre := nums[i]
			nums[i] = nums[i-1] + 1
			moved += nums[i] - pre
		}
	}
	return moved
}
