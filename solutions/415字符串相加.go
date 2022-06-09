package solutions

import "strconv"

func AddStrings(num1 string, num2 string) string {
	var res string
	n := max(len(num1), len(num2))
	// 进位
	var carry int64
	for i := 0; i < n; i++ {
		x, y := len(num1)-i-1, len(num2)-i-1
		var a, b int64
		if x < 0 {
			a = 0
		} else {
			a, _ = strconv.ParseInt(string(num1[x]), 10, 64)
		}
		if y < 0 {
			b = 0
		} else {
			b, _ = strconv.ParseInt(string(num2[y]), 10, 64)
		}
		cur := (a + b + carry) % 10
		res += strconv.FormatInt(cur, 10)
		carry = (a + b + carry) / 10
	}
	return reverseStr(res)
}

func reverseStr(s string) string {
	arr := []rune(s)
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
	return string(arr)
}
