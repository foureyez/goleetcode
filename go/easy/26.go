package easy

func removeDuplicates(nums []int) int {
	for i := 0; i < len(nums)-1; {
		if nums[i] != nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
			continue
		}
		i++
	}
	return len(nums)
}
