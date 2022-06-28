package leetcode

func sortedSquares(nums []int) []int {
	size := len(nums)
	pos := 0

	if size == 1 {
		one := make([]int, size)
		one[0] = nums[0] * nums[0]
		return one
	}

	for nums[pos] < 0 && pos < size {
		pos++
	}

	neg := pos - 1

	result := make([]int, size)
	var i int

	for neg >= 0 && size > pos {
		if nums[neg]*nums[neg] < nums[pos]*nums[pos] {
			result[i] = nums[neg] * nums[neg]
			neg--
		} else {
			result[i] = nums[pos] * nums[pos]
			pos++
		}
		i++
	}

	for neg >= 0 {
		result[i] = nums[neg] * nums[neg]
		neg--
		i++
	}

	for pos <= size-1 {
		result[i] = nums[pos] * nums[pos]
		pos++
		i++
	}

	return result
}
