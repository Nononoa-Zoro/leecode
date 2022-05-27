package 排序算法

func QuickSort(nums []int, left int, right int) {
	if left > right {
		return
	}
	pivot := nums[left]
	i, j := left, right
	for i != j {
		for i < j && nums[j] >= pivot {
			j--
		}
		for i < j && nums[i] <= pivot {
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	// i=j 将pivot的值放在i位置 left位置的值修改为i位置的值
	nums[left] = nums[i]
	nums[i] = pivot
	QuickSort(nums, left, i-1)
	QuickSort(nums, i+1, right)
}
