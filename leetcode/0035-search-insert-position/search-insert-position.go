package leetcode

func searchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for high >= low {
		mid := low + (high-low)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	// When high >= low and target is not in the list
	// Return closest ones left
	return low
}
