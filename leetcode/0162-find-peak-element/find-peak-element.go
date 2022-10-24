package leetcode

func findPeakElement(nums []int) int {
	// The array consist of sorted arrays -> may be reversed though
	// We need to return end or start of one of the sub-array
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] < nums[mid+1] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l // when l == r
}
