package solutions

var res [][]int

func Permute(nums []int) [][]int {
	if len(nums) == 0 {
		return res
	}
	solve(nums, 0, len(nums)-1)
	return res
}

func solve(nums []int, start, end int) {
	if start == end {
		var tmp []int
		for _, v := range nums {
			tmp = append(tmp, v)
		}
		res = append(res, tmp)
	}
	for i := start; i <= end; i++ {
		nums[i], nums[start] = nums[start], nums[i]
		solve(nums, start+1, end)
		nums[i], nums[start] = nums[start], nums[i]
	}
}
