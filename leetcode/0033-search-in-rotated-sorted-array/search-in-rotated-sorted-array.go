package leetcode

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	// One half of the nums should always be sorted in ascended order
	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] == target {
			return mid
		}

		if nums[l] <= nums[mid] {
			// if ascending order is in the left
			if target >= nums[l] && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			// if ascending order is in the right
			if target <= nums[r] && target > nums[mid] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return -1
}
