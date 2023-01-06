package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		diff := target - v
		if _, ok := m[diff]; ok {
			return []int{m[diff], i}
		}
		m[v] = i
	}
	return []int{-1, -1}
}
