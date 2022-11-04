package leetcode

func findMaxConsecutiveOnes2(nums []int) int {
	rightLen, leftLen, max := 0, 0, 0

	if len(nums) < 2 {
		return len(nums)
	}

	for _, v := range nums {
		if v == 1 {
			rightLen++
		} else {
			leftLen = rightLen
			rightLen = 0
		}

		if rightLen+leftLen+1 > max {
			max = rightLen + leftLen + 1
		}
	}

	return max
}
