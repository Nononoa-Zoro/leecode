package solutions

func LongestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	// 初始化dp dp[i][j]表示第i-j索引位置的子串是否为回文串
	dp := make([][]bool, n, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	// 初始化每个字符是一个回文串
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	maxLen := 1 // 回文串最大长度
	start := 0  // 记录结果最长子串的左边界索引
	// 枚举子串长度
	for L := 2; L <= n; L++ {
		for i := 0; i < n; i++ {
			j := L + i - 1
			// 右边界超出字符串长度
			if j >= n {
				break
			}

			// 如果左边界的字符!=右边界的字符
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				// 子串长度为2或者3的时候 且两边界的字符相等
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			// 更新最大子串长度和子串的起始位置
			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				start = i
			}
		}
	}
	return s[start : start+maxLen]
}
