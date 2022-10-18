package leetcode

func findMin(nums []int) int {
	l := 0
	r := len(nums) - 1

	if nums[l] < nums[r] {
		return nums[l]
	}

	for l+1 < r {
		mid := l + (r-l)/2
		if nums[mid] > nums[r] {
			l = mid
		} else {
			r = mid
		}
	}

	if nums[l] < nums[r] {
		return nums[l]
	}
	return nums[r]
}
