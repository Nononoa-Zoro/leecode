package 排序算法

func Insert(nums []int) {
	// 枚举n-1次
	// 第一个元素自身表示一个有序数组
	for i := 1; i < len(nums); i++ {
		j := i
		t := nums[j]
		for j >= 1 && t < nums[j-1] {
			nums[j] = nums[j-1]
			j--
		}
		nums[j] = t
	}
}
