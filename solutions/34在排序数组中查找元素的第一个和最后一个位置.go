package solutions

func searchRange(nums []int, target int) []int {
	var res []int
	index := binarySearch(nums, 0, len(nums)-1, target)
	if index == -1 {
		res = append(res, -1, -1)
		return res
	}
	left, right := index, index
	for left >= 0 && nums[left] == nums[index] {
		left--
	}
	if left < 0 {
		res = append(res, 0)
	} else {
		left += 1
		res = append(res, left)
	}

	for right < len(nums) && nums[right] == nums[index] {
		right++
	}
	if right == len(nums) {
		res = append(res, len(nums)-1)
	} else {
		right -= 1
		res = append(res, right)
	}
	return res
}
