package leetcode

func removeDuplicates(nums []int) int {
	// [1, 1, 2, 3, 4, 5, 5]
	//  ^  ^               | Iterators will move through the array
	// if nums[left] == nums[right] => remove nums[second] (Send to end of the list in this case)

	// [1, 1, 2, 3, 3]
	//  ^     ^ -> Right pointer will move and replace with [left + 1] and [right]

	// Base case
	if len(nums) <= 1 {
		return len(nums)
	}
	// Pointers
	var left, right = 0, 1

	for right < len(nums) {

		if nums[left] >= nums[right] {
			right++
		} else if nums[left] < nums[right] {
			nums[left+1], nums[right] = nums[right], nums[left+1]
			left++
		}

	}

	return left + 1
}
