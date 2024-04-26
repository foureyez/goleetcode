package easy

/*
https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/?envType=study-plan-v2&envId=leetcode-75
*/
func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandies := 0
	out := make([]bool, len(candies))
	for _, c := range candies {
		if c > maxCandies {
			maxCandies = c
		}
	}

	for i, c := range candies {
		if c+extraCandies >= maxCandies {
			out[i] = true
		}
	}
	return out
}
