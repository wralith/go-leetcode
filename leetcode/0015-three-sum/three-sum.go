package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	var left, right int
	res := [][]int{}
	visited := make(map[int]bool)
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if _, ok := visited[nums[i]]; ok {
			continue
		}

		left, right = i+1, len(nums)-1
		for left < right {
			sum := nums[left] + nums[right] + nums[i]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				// To capture situations like `Example 3` In test
				for nums[left] == nums[left-1] && left < right {
					left++
				}
			}
			if sum < 0 {
				left++
			}
			if sum > 0 {
				right--
			}
		}
		visited[nums[i]] = true
	}

	return res
}
