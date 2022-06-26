package solutions

// findMedianSortedArrays 在两个升序数组中找到中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	length := len1 + len2
	arr := make([]int, length, length)
	i, j, k := 0, 0, 0
	for i != len1 && j != len2 {
		if nums1[i] < nums2[j] {
			arr[k] = nums1[i]
			k++
			i++
		} else {
			arr[k] = nums2[j]
			k++
			j++
		}
	}
	for i != len1 {
		arr[k] = nums1[i]
		i++
		k++
	}
	for j != len2 {
		arr[k] = nums2[j]
		j++
		k++
	}
	middle := length >> 1
	if length&1 != 0 {
		return float64(arr[middle])
	} else {
		return float64(arr[middle-1]+arr[middle]) / 2
	}
}
