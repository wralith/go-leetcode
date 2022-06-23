package leetcode

func search(nums []int, target int) int {
	lowest := 0
	highest := len(nums) - 1

	for highest >= lowest {
		half := (lowest + highest) / 2
		if nums[half] == target {
			return half
		}
		if nums[half] > target {
			highest = half - 1
		}
		if nums[half] < target {
			lowest = half + 1
		}
	}

	return -1 // Target is not in the nums array
}
