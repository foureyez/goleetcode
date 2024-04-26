package easy

/*
https://leetcode.com/problems/longest-ideal-subsequence/?envType=daily-question&envId=2024-04-25
*/
func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	len := gcd(len(str1), len(str2))
	return str1[:len]
}

func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}
