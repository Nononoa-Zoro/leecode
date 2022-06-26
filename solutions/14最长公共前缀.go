package solutions

import "strings"

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

// longestCommonPrefix2 解法2
// 选择任意一个字符串作为prefix 遍历数组 如果不是子串 prefix往前截取
func longestCommonPrefix2(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	prefix := strs[0]
	for _, s := range strs {
		for strings.Index(s, prefix) != 0 {
			if prefix == "" {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}
