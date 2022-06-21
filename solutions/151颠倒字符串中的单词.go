package solutions

import "strings"

func reverseWords(s string) string {
	splitArr := strings.Split(s, " ")
	var words []string
	for _, s := range splitArr {
		if s == " " || s == "" {
			continue
		}
		words = append(words, s)
	}
	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-1-i] = words[len(words)-1-i], words[i]
	}
	var builder strings.Builder
	for i := 0; i < len(words); i++ {
		builder.WriteString(words[i])
		if i != len(words)-1 {
			builder.WriteString(" ")
		}
	}
	return builder.String()
}
