package leetcode

func findMaxConsecutiveOnes(nums []int) int {
	count, max := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		} else {
			count = 0
		}

		if count > max {
			max = count
		}
	}

	return max
}
