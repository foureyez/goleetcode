package easy

/*
https://leetcode.com/problems/reverse-vowels-of-a-string/description/?envType=study-plan-v2&envId=leetcode-75
*/
func reverseVowels(s string) string {
	i := 0
	j := len(s) - 1
	rs := []rune(s)
	for i < j {
		if isVowel(rs[i]) && isVowel(rs[j]) {
			tmp := rs[i]
			rs[i] = rs[j]
			rs[j] = tmp
			i++
			j--
		} else if isVowel(rs[i]) {
			j--
		} else {
			i++
		}
	}
	return string(rs)
}

func isVowel(c rune) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
		c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
}
