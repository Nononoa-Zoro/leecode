package solutions

// BinarySearch 二分查找
func BinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		center := (left + right) >> 1
		if nums[center] < target {
			left = center + 1
		} else if nums[center] > target {
			right = center - 1
		} else {
			return center
		}
	}
	return -1
}

func BinarySearchRecursive(nums []int, left, right, target int) int {
	// 左指针>右指针 还没找到 返回-1
	if left > right {
		return -1
	}
	center := (left + right) >> 1
	if nums[center] == target {
		return center
	} else if nums[center] > target {
		BinarySearchRecursive(nums, left, center-1, target)
	} else {
		BinarySearchRecursive(nums, center+1, right, target)
	}
	return -1
}
