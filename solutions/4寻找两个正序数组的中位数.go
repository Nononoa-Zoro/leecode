package solutions

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1)
	m := len(nums2)
	length := n + m
	// nums1 nums2的指针
	a, b := 0, 0
	// last保存每次循环的当前元素左边的那个元素
	last, cur := -1, -1
	for i := 0; i <= length/2; i++ {
		last = cur
		if a < n && (b >= m || nums1[a] < nums2[b]) {
			cur = nums1[a]
			a++
		} else {
			cur = nums2[b]
			b++
		}
	}
	// 总长度为偶数 选中间两个 取平均
	if length&1 == 0 {
		return float64(last+cur) / 2
	}
	return float64(cur)
}
