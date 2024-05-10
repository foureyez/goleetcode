package easy

func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums)-1; {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			continue
		}
		i++
	}
	return len(nums)
}
