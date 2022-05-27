package solutions

func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	if len(s) == 0 {
		return 0
	}
	var cur, res, right int
	for i := len(s) - 1; i >= 0; i-- {
		cur = m[s[i]]
		if cur < right {
			res -= cur
		} else {
			res += cur
		}
		right = cur
	}
	return res
}
