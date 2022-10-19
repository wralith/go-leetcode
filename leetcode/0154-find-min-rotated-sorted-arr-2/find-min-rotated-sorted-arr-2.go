package leetcode

func findMin2(nums []int) int {
	l := 0
	r := len(nums) - 1

	for l+1 < r {
		mid := l + (r-l)/2
		if nums[mid] > nums[r] {
			l++
		} else {
			r--
		}
	}

	if nums[l] < nums[r] {
		return nums[l]
	}
	return nums[r]
}
