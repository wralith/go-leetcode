package leetcode

func searchRange(nums []int, target int) []int {
	result := []int{-1, -1}
	if len(nums) == 0 {
		return result
	}

	result[0] = findLeft(nums, target)
	result[1] = findRight(nums, target)

	return result

}

func findLeft(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l+1 < r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid
		} else {
			r = mid
		}
	}
	if nums[l] == target {
		return l
	}
	if nums[r] == target {
		return r
	}
	return -1
}

func findRight(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l+1 < r {
		mid := l + (r-l)/2
		if nums[mid] <= target {
			l = mid
		} else {
			r = mid
		}
	}
	if nums[r] == target {
		return r
	}
	if nums[l] == target {
		return l
	}
	return -1
}
