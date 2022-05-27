package solutions

// BinarySearch 二分查找
func BinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		center := left + (right-left)/2
		if nums[center] < target {
			left = center + 1
		}
		if nums[center] > target {
			right = center - 1
		}
		return center
	}
	return -1
}
