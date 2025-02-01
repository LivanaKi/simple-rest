package concatenations

import (
	"strings"
)

func ConcatOne(str []string) string {
	result := ""
	for _, v := range str {
		result += v
	}
	return result
}

func ConcatTwo(str []string) string {
	return strings.Join(str, "")
}

func ConcatThree(str []string) string {
	var builder strings.Builder
	for _, v := range str {
		builder.WriteString(v)
	}
	return builder.String()
}
