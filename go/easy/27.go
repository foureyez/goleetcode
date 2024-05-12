package easy

func removeElement(nums []int, val int) int {
	var i, j int
	for ; i < len(nums); i++ {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}

// func removeElement(nums []int, val int) int {
// 	for i := 0; i < len(nums)-1; {
// 		if nums[i] == val {
// 			nums = append(nums[:i], nums[i+1:]...)
// 			continue
// 		}
// 		i++
// 	}
// 	return len(nums)
// }
