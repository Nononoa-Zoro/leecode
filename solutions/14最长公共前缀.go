package solutions

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	var res string
	minLen := minLengthOfStr(strs)
	for i := 0; i < minLen; i++ {
		// 第一个字符串当做模板
		template := strs[0][i]
		// 判断所有字符串的第i个位置元素是否相等
		for j := 0; j < len(strs); j++ {
			if strs[j][i] != template {
				return res
			}
		}
		res = res + string(template)
	}
	return res
}

// minLengthOfStr 返回数组中最短字符串的长度
func minLengthOfStr(strs []string) int {
	res := len(strs[0])
	for _, s := range strs {
		if len(s) < res {
			res = len(s)
		}
	}
	return res
}
