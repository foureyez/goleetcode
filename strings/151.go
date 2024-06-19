package strings

import (
	"strings"
)

/*
https://leetcode.com/problems/reverse-words-in-a-string/description/?envType=study-plan-v2&envId=leetcode-75
*/
func reverseWords(s string) string {
	var out strings.Builder
	splits := strings.Split(s, " ")
	for i := len(splits) - 1; i >= 0; i-- {
		if splits[i] == "" {
			continue
		}
		out.WriteString(splits[i] + " ")
	}

	return strings.Trim(out.String(), " ")
}
