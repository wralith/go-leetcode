package leetcode

func sortedSquares(nums []int) []int {
	size := len(nums)
	pos := 0

	// Only one item
	result := make([]int, size)
	if size == 1 {
		result[0] = nums[0] * nums[0]
		return result
	}

	for nums[pos] < 0 && pos < size-1 {
		pos++
	}

	neg := pos - 1

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

	// First and Last
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
