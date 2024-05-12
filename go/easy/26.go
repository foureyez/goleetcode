package easy

func removeDuplicates(nums []int) int {
	unique := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[unique] = nums[i]
			unique++
		}
	}
	return unique
}

// func removeDuplicates(nums []int) int {
// 	for i := 0; i < len(nums)-1; {
// 		if nums[i] != nums[i+1] {
// 			nums = append(nums[:i], nums[i+1:]...)
// 			continue
// 		}
// 		i++
// 	}
// 	return len(nums)
// }
