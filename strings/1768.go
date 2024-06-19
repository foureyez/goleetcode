package strings

import (
	"strings"
)

/*
*
https://leetcode.com/problems/merge-strings-alternately/?envType=study-plan-v2&envId=leetcode-75
*
*/
func mergeAlternately(word1 string, word2 string) string {
	w1 := len(word1)
	w2 := len(word2)
	minLen := min(w1, w2)

	var out strings.Builder
	for i := 0; i < minLen; i++ {
		out.WriteByte(word1[i])
		out.WriteByte(word2[i])
	}

	// NO need to check for this condition
	// Since whichever string is shorter will write nothing to output string
	// if w1 != w2 {
	// 	if w1 < w2 {
	// 		out.WriteString(word2[minLen:])
	// 	} else {
	// 		out.WriteString(word1[minLen:])
	// 	}
	// }

	out.WriteString(word2[minLen:])
	out.WriteString(word1[minLen:])

	return out.String()
}
