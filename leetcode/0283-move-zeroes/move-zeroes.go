package leetcode

func moveZeroes(nums []int) {
	// Pointer will stay if it finds zero
	// When number rather than zero found, will replace it with zero
	ptr := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if ptr != i {
				nums[i], nums[ptr] = nums[ptr], nums[i]
			}
			ptr++
		}
	}
}

// First Try
// Pop all the zeroes out of the array and keep count
// Then appends those zeroes to the end of the array
// func moveZeroes(nums []int) {
// 	zeroCount := 0
// 	i := 0
// 	for i < len(nums) {
// 		if nums[i] == 0 {
// 			zeroCount++
// 			nums = append(nums[:i], nums[i+1:]...)
// 		} else {
// 			i++
// 		}
// 	}

// 	for zeroCount != 0 {
// 		nums = append(nums, 0) // There should be better way :(
// 		zeroCount--
// 	}
// }
