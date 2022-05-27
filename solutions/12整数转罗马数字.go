package solutions

// https://leetcode-cn.com/problems/integer-to-roman/
func intToRoman(num int) string {
	const size = 13
	nums := [size]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	roman := [size]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	index := 0
	var s string
	for index < size {
		for num >= nums[index] {
			s += roman[index]
			num -= nums[index]
		}
		index++
	}
	return s
}
