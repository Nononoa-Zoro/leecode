package 排序算法

// Bubble 冒泡排序
// 比较n-1轮，每轮确定一个数的最终位置
// 第一轮：比较n-1次
// 第二轮：比较n-2次
// 第三轮：比较n-3次
// ...
// 第n-1轮：比较1次
// 算法复杂度 o(n^2)
func Bubble(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j-1] > nums[j] {
				swap(j-1, j, nums)
			}
		}
	}
}

// BubbleForSeq 有序数组冒泡排序
func BubbleForSeq(nums []int) {
	flag := true // 这次遍历是否发生交换 如果全局已经有序 不再进行之后的排序操作
	length := len(nums)
	for flag {
		flag = false
		for j := 0; j < length-1; j++ {
			if nums[j] > nums[j+1] {
				swap(j, j+1, nums)
				flag = true
			}
		}
		length--
	}
}
func swap(i, j int, nums []int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}
