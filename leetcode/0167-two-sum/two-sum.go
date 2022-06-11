package leetcode

func twoSum(numbers []int, target int) []int {

	first := 0               // Points left
	last := len(numbers) - 1 // Points right

	for first < last {
		sum := numbers[first] + numbers[last]
		if sum == target {
			return []int{first + 1, last + 1}
		}
		if sum < target {
			first++
		}
		if sum > target {
			last--
		}
	}

	// No solution!
	return []int{-1, -1}
}
