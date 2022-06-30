package solutions

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	for i := 0; i < len(v1) || i < len(v2); i++ {
		var x, y int64
		if i < len(v1) {
			x, _ = strconv.ParseInt(v1[i], 10, 64)
		}
		if i < len(v2) {
			y, _ = strconv.ParseInt(v2[i], 10, 64)
		}
		if x > y {
			return 1
		}
		if x < y {
			return -1
		}
	}
	return 0
}
