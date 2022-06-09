package solutions

func merge(nums1 []int, m int, nums2 []int, n int) {
	// 双指针从后往前遍历
	i, j, k := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 {
		if nums1[i] < nums2[j] {
			nums1[k] = nums2[j]
			k--
			j--
		} else {
			nums1[k] = nums1[i]
			i--
			k--
		}
	}
	// nums2数组还有元素没有遍历完成，全部append到nums1的最前面
	for j >= 0 {
		nums1[k] = nums2[j]
		k--
		j--
	}
}
