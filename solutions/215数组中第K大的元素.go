package solutions

func findKthLargest(nums []int, k int) int {
	return quickSort(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSort(nums []int, left, right, k int) int {
	if left == right {
		return nums[left]
	}
	pivotIndex := left
	pivotIndex = partition(nums, left, right, pivotIndex)
	if k == pivotIndex {
		return nums[k]
	} else if k < pivotIndex {
		return quickSort(nums, left, pivotIndex-1, k)
	}
	return quickSort(nums, pivotIndex+1, right, k)
}

func partition(nums []int, left, right, pivotIndex int) int {
	pivot := nums[pivotIndex]
	nums[pivotIndex], nums[right] = nums[right], nums[pivotIndex]
	storeIndex := left

	for i := left; i <= right; i++ {
		if nums[i] < pivot {
			nums[i], nums[storeIndex] = nums[storeIndex], nums[i]
			storeIndex++
		}
	}

	nums[storeIndex], nums[right] = nums[right], nums[storeIndex]
	return storeIndex
}
