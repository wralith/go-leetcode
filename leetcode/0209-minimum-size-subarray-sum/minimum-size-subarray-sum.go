package leetcode

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	l, r, sum, min := 0, 0, 0, n+1

	for r < n {
		sum += nums[r]
		r++
		for target <= sum {
			sum -= nums[l]
			l++
			windowLen := r - l + 1
			if windowLen < min {
				min = windowLen
			}

		}
	}

	// if min == n+1 {
	// 	return 0
	// }
	// return min
	return min % (n + 1)
}
