package solutions

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]bool)
	n := len(s)
	right, ans := 0, 0
	for i := 0; i < n; i++ {
		// 左指针往右移动一个位置 删除i之前字符的记录
		// 相当于左指针向右移动
		if i != 0 {
			delete(m, s[i-1])
		}
		// 右指针 如果不在map中 则继续往右移动
		for right < n && !m[s[right]] {
			m[s[right]] = true
			right++
		}
		// 这里退出循环 m[right]=true right位置对应的字符肯定是之前出现过的
		ans = max(ans, right-i)
	}
	return ans
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
