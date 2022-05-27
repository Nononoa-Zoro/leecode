package solutions

import (
	"fmt"
	"strings"
)

func Convert(s string, numRows int) string {
	table := table(len(s), numRows)
	fmt.Println(len(table))
	arr := make([][]rune, numRows, numRows)
	for i := 0; i < len(s); i++ {
		arr[table[i]] = append(arr[table[i]], charAt(s, i))
	}

	var builder strings.Builder
	for _, row := range arr {
		for _, i := range row {
			builder.WriteString(string(i))
		}
	}
	return builder.String()
}

func charAt(s string, i int) rune {
	return []rune(s)[i]
}

func table(len, numRows int) []int {
	table := make([]int, 0, len)
	itr := 0
	var flag bool
	for i := 0; i < len; i++ {
		if !flag && itr < numRows {
			table = append(table, itr)
			itr++
		}
		if itr == numRows {
			itr--
			flag = !flag
		}
		if flag && itr > 0 {
			itr--
			table = append(table, itr)
		}
		if itr == 0 {
			itr++
			flag = !flag
		}
	}
	return table
}
