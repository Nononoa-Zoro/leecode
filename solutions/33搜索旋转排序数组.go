package solutions

func Search(nums []int, target int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			res := binarySearch(nums, target, 0, i-1)
			if res != -1 {
				return res
			}
			res = binarySearch(nums, target, i, len(nums)-1)
			if res != -1 {
				return res
			}
			return -1
		}
	}

	return binarySearch(nums, target, 0, len(nums))
}

func binarySearch(nums []int, target int, left, right int) int {
	if target < nums[left] || target > nums[right] {
		return -1
	}
	for left <= right {
		center := left + (right-left)/2
		if nums[center] == target {
			return center
		} else if nums[center] > target {
			right = center - 1
		} else {
			left = center + 1
		}
	}
	return -1
}

// searchR 递归查找
func searchR(nums []int, left, right, target int) int {
	if left > right {
		return -1
	}
	center := (left + right) >> 1
	if nums[center] == target {
		return center
	}
	if nums[center] > nums[right] {
		if target > nums[left] && target < nums[center] {
			return searchR(nums, left, center-1, target)
		}
		return searchR(nums, center+1, right, target)
	} else {
		if target > nums[center] && target < nums[right] {
			return searchR(nums, center+1, right, target)
		}
		return searchR(nums, left, center-1, target)
	}
}
