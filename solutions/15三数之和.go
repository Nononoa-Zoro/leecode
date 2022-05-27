package solutions

import "sort"

// https://leetcode-cn.com/problems/3sum/submissions/
func threeSum(nums []int) [][]int {
	var res [][]int
	if len(nums) < 3 {
		return res
	}
	sort.Ints(nums)
	if nums[0] > 0 {
		return res
	}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		L, R := i+1, len(nums)-1
		for L < R {
			sum := nums[i] + nums[L] + nums[R]
			if sum == 0 {
				row := []int{nums[i], nums[L], nums[R]}
				res = append(res, row)
				for L < R && nums[L] == nums[L+1] {
					L++
				}
				L++
				for L < R && nums[R] == nums[R-1] {
					R--
				}
				R--
			} else if sum < 0 {
				L++
			} else {
				R--
			}
		}
	}
	return res
}
