package medium

/*
https://leetcode.com/problems/product-of-array-except-self/description/?envType=study-plan-v2&envId=leetcode-75
*/
func productExceptSelf(nums []int) []int {
	n := len(nums)
	prefix := make([]int, n)
	suffix := make([]int, n)

	prefix[0] = nums[0]
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] * nums[i]
	}

	suffix[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i]
	}

	for i := 0; i < n; i++ {
		switch i {
		case 0:
			nums[i] = suffix[1]
		case n - 1:
			nums[i] = prefix[n-2]
		default:
			nums[i] = prefix[i-1] * suffix[i+1]
		}
	}
	return nums
}
