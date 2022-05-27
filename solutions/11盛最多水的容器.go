package solutions

// https://leetcode-cn.com/problems/container-with-most-water/
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	ans := -1
	for left != right {
		cur := min(height[left], height[right]) * (right - left)
		ans = max(ans, cur)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
